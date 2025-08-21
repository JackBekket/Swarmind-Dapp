import React, { useState } from 'react';
import { ethers } from 'ethers';
import {
  swmAddress,
  poolAddress,
  chainId,
  chainName,
  curName,
  curSymbol,
  chainRPC,
  blockExplorer,
  permitDeadlineSeconds,
} from '../constants';
import { erc20Abi, poolAbi } from '../abis';

const SLEEP = (ms: number) => new Promise(res => setTimeout(res, ms));

const isTransientError = (err: any) => {
  if (!err) return false;
  const msg = (err.reason || err.message || '').toString().toLowerCase();
  const code = err.code?.toString?.().toLowerCase() || '';
  // common transient patterns
  if (/missing revert data/i.test(msg)) return true;
  if (/connection.*reset|econnreset|etimedout|timeout|network error|503|502|econnrefused|socket hang up/i.test(msg)) return true;
  if (code === 'call_exception' || code === 'servererror' || code === 'networkerror') return true;
  return false;
};

async function retryWrapper<T>(fn: () => Promise<T>, attempts = 5, baseDelay = 300): Promise<T> {
  let lastErr: any;
  for (let i = 0; i < attempts; i++) {
    try {
      return await fn();
    } catch (err: any) {
      lastErr = err;
      if (!isTransientError(err) || i === attempts - 1) break;
      const jitter = Math.floor(Math.random() * 200);
      const delay = Math.min(baseDelay * 2 ** i + jitter, 5000);
      // small log for debugging
      console.warn(`Transient RPC error, retrying ${i + 1}/${attempts} after ${delay}ms:`, err.message || err);
      await SLEEP(delay);
    }
  }
  throw lastErr;
}

const Deposit: React.FC = () => {
  const [transferFromAmount, setTransferFromAmount] = useState<number>(0);
  const [depositAmount, setDepositAmount] = useState<number>(0);
  const [isConnected, setIsConnected] = useState<boolean>(false);

  const connectWalletAndSwitchNetwork = async () => {
    if (!window.ethereum) {
      alert('MetaMask is not installed. Redirecting...');
      window.location.href = 'https://metamask.io/download.html';
      return;
    }

    const networkParams = {
      chainId: ethers.toBeHex(chainId),
      chainName,
      nativeCurrency: { name: curName, symbol: curSymbol, decimals: 18 },
      rpcUrls: [chainRPC],
      blockExplorerUrls: [blockExplorer],
    };

    try {
      await window.ethereum.request({ method: 'eth_requestAccounts' });

      try {
        await window.ethereum.request({
          method: 'wallet_switchEthereumChain',
          params: [{ chainId: ethers.toBeHex(chainId) }],
        });
      } catch (switchError: any) {
        if (switchError.code === 4902) {
          await window.ethereum.request({
            method: 'wallet_addEthereumChain',
            params: [networkParams],
          });
        }
      }

      setIsConnected(true);
      alert('MetaMask connected and network switched!');
    } catch (error) {
      console.error('Connection or network switch error:', error);
      alert('Не удалось подключиться/поменять сеть' );
    }
  };

  const checkNetwork = async (): Promise<boolean> => {
    if (!window.ethereum) return false;
    const provider = new ethers.BrowserProvider(window.ethereum);
    const network = await retryWrapper(() => provider.getNetwork());
    // network.chainId is BigInt in ethers v6
    return BigInt(network.chainId) === BigInt(chainId);
  };

  // helper to wait for tx safely (timeout + final receipt check)
  async function waitForTx(provider: ethers.BrowserProvider, txHash: string, confirmations = 1, timeoutMs = 120000) {
    try {
      // try provider.waitForTransaction with timeout
      const receipt = await provider.waitForTransaction(txHash, confirmations, timeoutMs);
      if (receipt) return receipt;
    } catch (err) {
      // ignore, we'll attempt a direct receipt check below
      console.warn('waitForTransaction failed or timed out, checking receipt directly...', err);
    }

    // As fallback, poll for receipt a few times
    const pollInterval = 2000;
    const maxTries = Math.ceil(timeoutMs / pollInterval);
    for (let i = 0; i < maxTries; i++) {
      const r = await provider.getTransactionReceipt(txHash);
      if (r && r.blockNumber) return r;
      await SLEEP(pollInterval);
    }
    throw new Error('Timeout waiting for tx receipt');
  }

  const gaslessApproveWithPermit = async () => {
    if (!window.ethereum) return alert('MetaMask not found');

    if (!await checkNetwork()) {
      alert(`Please switch to ${chainName} network`);
      return;
    }

    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      const owner = await signer.getAddress();
      const contract = new ethers.Contract(swmAddress, erc20Abi, signer);

      // Check if contract supports permit (with retries for flaky RPC)
      try {
        await retryWrapper(() => contract.nonces(owner), 4, 200);
      } catch (e) {
        alert('This token does not support permit functionality or the node is not responding to nonces()');
        console.error('nonces check failed:', e);
        return;
      }

      // read nonce & name with retry wrapper
      const rawNonce = await retryWrapper(() => contract.nonces(owner), 4, 200);
      // rawNonce might be BigNumber -> toBigInt
      const nonce = BigInt(rawNonce.toString?.() ?? rawNonce);
      const deadline = Math.floor(Date.now() / 1000) + permitDeadlineSeconds;
      const value = ethers.parseUnits(transferFromAmount.toString(), 9);

      // domain.name may call name() on token; wrap with retry
      const tokenName = await retryWrapper(() => contract.name(), 4, 200);

      const domain = {
        name: tokenName,
        version: '1',
        chainId: Number(chainId), // ethers v6 accepts number for signTypedData domain
        verifyingContract: swmAddress,
      };

      const types = {
        Permit: [
          { name: 'owner', type: 'address' },
          { name: 'spender', type: 'address' },
          { name: 'value', type: 'uint256' },
          { name: 'nonce', type: 'uint256' },
          { name: 'deadline', type: 'uint256' },
        ],
      };

      const message = {
        owner,
        spender: poolAddress,
        value,
        nonce,
        deadline,
      };

      // signTypedData can also fail intermittently; retry it
      const signature = await retryWrapper(() => signer.signTypedData(domain, types, message), 4, 200);
      const { v, r, s } = ethers.Signature.from(signature);

      // perform permit tx with retries for transient RPC errors
      let tx: any;
      try {
        tx = await retryWrapper(
          () => contract.permit(owner, poolAddress, value, deadline, v, r, s, { gasLimit: 1000000 }),
          3,
          400,
        );
      } catch (err: any) {
        console.error('Failed to send permit transaction after retries:', err);
        // If it's not transient - show it; if it was transient and still failed, let user know
        alert('Permit transaction failed (node unstable). Попробуйте ещё раз.');
        return;
      }

      // wait for confirmation robustly
      try {
        await waitForTx(provider, tx.hash, 1, 120000);
      } catch (err) {
        // If provider wait timed out, check receipt once more
        console.warn('Waiting for permit tx failed:', err);
        const receipt = await provider.getTransactionReceipt(tx.hash);
        if (!receipt || !receipt.blockNumber) {
          alert('Permit tx not confirmed (timeout). Проверьте в обозревателе и повторите при необходимости.');
          return;
        }
      }

      alert('Gasless approve (permit) successful!');
    } catch (error: any) {
      console.error('Permit error:', error);
      const userMsg = (error?.reason || error?.message || 'Unknown error').toString();
      alert(`Permit failed: ${userMsg}`);
    }
  };

  const checkAllowanceAndBalance = async (): Promise<{ hasAllowance: boolean; hasBalance: boolean }> => {
    const provider = new ethers.BrowserProvider(window.ethereum);
    const signer = await provider.getSigner();
    const owner = await signer.getAddress();

    const tokenContract = new ethers.Contract(swmAddress, erc20Abi, signer);
    // it's safe to read balance & allowance with retries
    const balance = await retryWrapper(() => tokenContract.balanceOf(owner), 4, 200);
    const allowance = await retryWrapper(() => tokenContract.allowance(owner, poolAddress), 4, 200);
    const value = ethers.parseUnits(depositAmount.toString(), 9);

    // convert to comparable BigInt if necessary
    const balanceBig = BigInt(balance.toString?.() ?? balance);
    const allowanceBig = BigInt(allowance.toString?.() ?? allowance);

    return {
      hasAllowance: allowanceBig >= BigInt(value.toString?.() ?? value),
      hasBalance: balanceBig >= BigInt(value.toString?.() ?? value),
    };
  };

  const depositCredit = async () => {
    if (!window.ethereum) return alert('MetaMask is not installed');

    if (!await checkNetwork()) {
      alert(`Please switch to ${chainName} network`);
      return;
    }

    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      const pool = new ethers.Contract(poolAddress, poolAbi, signer);
      const value = ethers.parseUnits(depositAmount.toString(), 9);

      // Check if function exists in ABI
      const functionExists = pool.interface.hasFunction('DepositCredit');
      if (!functionExists) {
        alert('DepositCredit function not found in contract ABI');
        return;
      }

   
      const { hasAllowance, hasBalance } = await checkAllowanceAndBalance();

      if (!hasBalance) {
        alert('Insufficient token balance');
        return;
      }

      if (!hasAllowance) {
        alert('Insufficient allowance. Please approve tokens first.');
        return;
      }

      // Encode function data manually to ensure it's correct
      const data = pool.interface.encodeFunctionData('DepositCredit', [value]);

      // send tx with retry on transient errors
      let tx: any;
      try {
        tx = await retryWrapper(
          () => signer.sendTransaction({ to: poolAddress, data, gasLimit: 1000000 }),
          3,
          400,
        );
      } catch (err: any) {
        console.error('Failed to send DepositCredit tx:', err);
        alert('Не удалось отправить DepositCredit (RPC unstable). Попробуйте ещё раз.');
        return;
      }

      // wait robustly
      try {
        await waitForTx(provider, tx.hash, 1, 120000);
      } catch (err) {
        console.warn('Deposit tx wait failed:', err);
        const receipt = await provider.getTransactionReceipt(tx.hash);
        if (!receipt || !receipt.blockNumber) {
          alert('DepositCredit tx not confirmed (timeout). Проверьте обозреватель.');
          return;
        }
      }

      alert('DepositCredit successful!');
    } catch (error: any) {
      console.error('Deposit error:', error);
      alert(`DepositCredit failed: ${error?.reason || error?.message || error}`);
    }
  };

  return (
    <div className="p-4 space-y-6">
      <h1 className="text-xl font-bold">SWM Deposit DApp</h1>

      {isConnected ? (
        <div className="px-4 py-2 rounded shadow bg-green-500 text-white inline-block">
          Connected!
        </div>
      ) : (
        <button
          onClick={connectWalletAndSwitchNetwork}
          className="px-4 py-2 rounded shadow"
        >
          Connect & Switch Network
        </button>
      )}

      <section className="space-y-2">
        <h3 className="font-semibold">Approve + TransferFrom (Gasless)</h3>
        <input
          type="number"
          value={transferFromAmount}
          onChange={(e) => setTransferFromAmount(Number(e.target.value))}
          placeholder="Amount"
          className="border p-2 rounded"
          min="0"
          step="0.001"
        />
        <div className="flex space-x-2">
          <button
            onClick={gaslessApproveWithPermit}
            className="px-4 py-2 rounded shadow"
            disabled={transferFromAmount <= 0}
          >
            Permit (Gasless)
          </button>
        </div>
      </section>

      <section className="space-y-2">
        <h3 className="font-semibold">Deposit Credit</h3>
        <input
          type="number"
          value={depositAmount}
          onChange={(e) => setDepositAmount(Number(e.target.value))}
          placeholder="Amount"
          className="border p-2 rounded"
          min="0"
          step="0.001"
        />
        <button
          onClick={depositCredit}
          className="px-4 py-2 rounded shadow"
          disabled={depositAmount <= 0}
        >
          DepositCredit
        </button>
      </section>
    </div>
  );
};

export default Deposit;
