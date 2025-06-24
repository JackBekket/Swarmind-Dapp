import React, { useState } from 'react';
import { ethers } from 'ethers';
import {
  receiverAddress,
  swmAddress,
  chainId,
  chainName,
  curName,
  curSymbol,
  curDec,
  chainRPC,
  blockExplorer
} from '../constants';
import { erc20Abi } from '../abis';

declare global {
  interface Window {
    ethereum: any;
  }
}

const Deposit: React.FC = () => {
  const [amount, setAmount] = useState<number>(0);

  const connectWalletAndSwitchNetwork = async () => {
    if (!window.ethereum) {
      alert('MetaMask is not installed. Redirecting...');
      window.location.href = 'https://metamask.io/download.html';
      return;
    }
  
  
     const networkParams = {
      chainId: chainId,
      chainName: chainName,
      nativeCurrency: {
        name: curName,
        symbol: curSymbol,
        decimals: curDec,
      },
      rpcUrls: [chainRPC],
      blockExplorerUrls: [blockExplorer],
    }; 
  
    try {
      await window.ethereum.request({ method: 'eth_requestAccounts' });

     await window.ethereum.request({
        method: 'wallet_addEthereumChain',
        params: [networkParams],
      });
   
      await window.ethereum.request({
        method: 'wallet_switchEthereumChain',
        params: [{ chainId }],
      });
  
      alert('MetaMask connected and network switched!');
    } catch (error) {
      console.error('Connection or network switch error:', error);
    }
  };
  


  const transferSWM = async () => {
    if (!window.ethereum) {
      alert('MetaMask is not installed');
      return;
    }

    try {
      const provider = new ethers.BrowserProvider(window.ethereum);
      await provider.send('eth_requestAccounts', []);
      const signer = await provider.getSigner();

      const swmContract = new ethers.Contract(swmAddress, erc20Abi, signer);
      const amountInWei = ethers.parseUnits(amount.toString(), 9); // SWM uses 9 decimals

      const tx = await swmContract.transfer(receiverAddress, amountInWei);
      await tx.wait();

      alert('Transfer successful!');
    } catch (error: any) {
      if (error.code === 4001) {
        alert('Transaction was rejected by the user.');
      } else {
        console.error(error);
        alert('An error occurred during the transfer.');
      }
    }
  };

  return (
    <div>
      <h1>SWM Deposit DApp</h1>

      <button onClick={connectWalletAndSwitchNetwork}>
        Connect MetaMask & Switch Network
      </button>

      <input
        type="number"
        value={amount}
        onChange={(e) => setAmount(Number(e.target.value))}
        placeholder="Amount"
      />

      <button onClick={transferSWM}>Transfer SWM</button>
    </div>
  );
};

export default Deposit;