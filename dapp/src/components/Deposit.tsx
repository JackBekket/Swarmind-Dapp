import React, { useState } from 'react';
import { ethers } from 'ethers';
import {
  receiverAddress,
  swmAddress,
  poolAddress,
  chainId,
  chainName,
  curName,
  curSymbol,
  curDec,
  chainRPC,
  blockExplorer,
  permitDeadlineSeconds,
} from '../constants';
import { erc20Abi, poolAbi } from '../abis';



const Deposit: React.FC = () => {
  const [amount, setAmount] = useState<number>(0);
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
      chainId,
      chainName,
      nativeCurrency: { name: curName, symbol: curSymbol, decimals: curDec },
      rpcUrls: [chainRPC],
      blockExplorerUrls: [blockExplorer],
    };

    try {
      await window.ethereum.request({ method: 'eth_requestAccounts' });
      await window.ethereum.request({ method: 'wallet_addEthereumChain', params: [networkParams] });
      await window.ethereum.request({ method: 'wallet_switchEthereumChain', params: [{ chainId }] });
      setIsConnected(true);
      alert('MetaMask connected and network switched!');
    } catch (error) {
      console.error('Connection or network switch error:', error);
    }
  };

  const transferSWM = async () => {
    if (!window.ethereum) return alert('MetaMask is not installed');
    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      await provider.send('eth_requestAccounts', []);
      const signer = await provider.getSigner();
      const swmContract = new ethers.Contract(swmAddress, erc20Abi, signer);
      const value = ethers.parseUnits(amount.toString(), 9);
      const tx = await swmContract.transfer(receiverAddress, value);
      await tx.wait();
      alert('Transfer successful!');
    } catch (error: any) {
      console.error(error);
      alert(error.code === 4001 ? 'Transaction rejected.' : 'Error during transfer.');
    }
  };

  const gaslessApproveWithPermit = async () => {
    if (!window.ethereum) return alert('MetaMask not found');
    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      const owner = await signer.getAddress();
      const contract = new ethers.Contract(swmAddress, erc20Abi, signer);
      const nonce = await contract.nonces(owner);
      const deadline = Math.floor(Date.now() / 1000) + permitDeadlineSeconds;
      const value = ethers.parseUnits(transferFromAmount.toString(), 9);
      const { chainId: currentChain } = await provider.getNetwork();

      const domain = {
        name: await contract.name(),
        version: '1',
        chainId: currentChain,
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
      const message = { owner, spender: poolAddress, value: value.toString(), nonce: nonce.toString(), deadline };

      const signature = await signer.signTypedData(domain, types, message);
      const sigObj = ethers.Signature.from(signature);
      const { v, r, s } = sigObj;

      const tx = await contract.permit(owner, poolAddress, value, deadline, v, r, s);
      await tx.wait();
      alert('Gasless approve (permit) successful!');
    } catch (error) {
      console.error(error);
      alert('Permit failed.');
    }
  };

  //TODO: do we actually need this?
  const transferFromSWM = async () => {
    if (!window.ethereum) return alert('MetaMask is not installed');
    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      const swmContract = new ethers.Contract(swmAddress, erc20Abi, signer);
      const from = await signer.getAddress();
      const value = ethers.parseUnits(transferFromAmount.toString(), 9);
      const tx = await swmContract.transferFrom(from, receiverAddress, value);
      await tx.wait();
      alert('transferFrom successful!');
    } catch (error) {
      console.error(error);
      alert('transferFrom failed.');
    }
  };

  const depositCredit = async () => {
    if (!window.ethereum) return alert('MetaMask is not installed');
    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      const signer = await provider.getSigner();
      const pool = new ethers.Contract(poolAddress, poolAbi, signer);
      const value = ethers.parseUnits(depositAmount.toString(), 9);
      const tx = await pool.DepositCredit(value);
      await tx.wait();
      alert('DepositCredit successful!');
    } catch (error) {
      console.error(error);
      alert('DepositCredit failed.');
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
  <button onClick={connectWalletAndSwitchNetwork} className="px-4 py-2 rounded shadow">
    Connect & Switch Network
  </button>
)}

      <section className="space-y-2">
        <h3 className="font-semibold">Transfer SWM</h3>
        <input
          type="number"
          value={amount}
          onChange={(e) => setAmount(Number(e.target.value))}
          placeholder="Amount"
          className="border p-2 rounded"
        />
        <button onClick={transferSWM} className="px-4 py-2 rounded shadow">
          Transfer
        </button>
      </section>

      <section className="space-y-2">
        <h3 className="font-semibold">Approve + TransferFrom (Gasless)</h3>
        <input
          type="number"
          value={transferFromAmount}
          onChange={(e) => setTransferFromAmount(Number(e.target.value))}
          placeholder="Amount"
          className="border p-2 rounded"
        />
        <div className="flex space-x-2">
          <button onClick={gaslessApproveWithPermit} className="px-4 py-2 rounded shadow">
            Permit (Gasless)
          </button>
          <button
  onClick={transferFromSWM}
  className="px-4 py-2 rounded shadow bg-gray-400 cursor-not-allowed"
  disabled
>
  TransferFrom
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
        />
        <button onClick={depositCredit} className="px-4 py-2 rounded shadow">
          DepositCredit
        </button>
      </section>
    </div>
  );
};

export default Deposit;
