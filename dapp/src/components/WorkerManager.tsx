import React, { useState } from 'react';
import { ethers } from 'ethers';
import {
  poolAddress,
  chainId,
  chainName,
  curName,
  curSymbol,
  curDec,
  chainRPC,
  blockExplorer,
} from '../constants';
import { poolAbi } from '../abis';

interface WorkerInfo {
  wallet: string;
  isApproved: boolean;
  isBanned: boolean;
}

const WorkerManager: React.FC = () => {
  // Connection
  const [isConnected, setIsConnected] = useState(false);
  const provider = typeof window !== 'undefined' && (window as any).ethereum ? new ethers.BrowserProvider((window as any).ethereum) : null;

  // Inputs
  const [workerKey, setWorkerKey] = useState('');
  const [newWorkerKey, setNewWorkerKey] = useState('');

  // Display
  const [workerInfo, setWorkerInfo] = useState<WorkerInfo | null>(null);

  const connectWalletAndSwitchNetwork = async () => {
    if (!provider) {
      alert('MetaMask is not installed.');
      return;
    }

    const networkParams = {
      chainId,
      chainName,
      nativeCurrency: { name: curName, symbol: curSymbol, decimals: curDec },
      rpcUrls: [chainRPC],
      blockExplorerUrls: [blockExplorer],
    } as any;

    try {
      await provider!.send('eth_requestAccounts', []);
      await provider!.send('wallet_addEthereumChain', [networkParams]);
      await provider!.send('wallet_switchEthereumChain', [{ chainId }]);
      setIsConnected(true);
    } catch (err) {
      console.error(err);
      alert('Failed to connect or switch network');
    }
  };

  const getSignerPool = async () => {
    if (!provider) throw new Error('No provider');
    const signer = await provider.getSigner();
    return new ethers.Contract(poolAddress, poolAbi, signer);
  };

  const getReadPool = () => {
    if (!provider) throw new Error('No provider');
    return new ethers.Contract(poolAddress, poolAbi, provider);
  };

  // CRUD operations
  const registerWorker = async () => {
    try {
      const pool = await getSignerPool();
      const tx = await pool.RegisterWorker(workerKey);
      await tx.wait();
      alert('Worker registered');
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error registering worker');
    }
  };

  const approveWorker = async () => {
    try {
      const pool = await getSignerPool();
      const tx = await pool.ApproveWorker(workerKey);
      await tx.wait();
      alert('Worker approved');
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error approving worker');
    }
  };

  const unregisterWorker = async () => {
    try {
      const pool = await getSignerPool();
      const tx = await pool.Unregister(workerKey);
      await tx.wait();
      alert('Worker unregistered');
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error unregistering worker');
    }
  };

  const updateWorker = async () => {
    try {
      const pool = await getSignerPool();
      const tx = await pool.UpdateWorker(workerKey, newWorkerKey);
      await tx.wait();
      alert('Worker key updated');
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error updating worker');
    }
  };

  const banWorker = async () => {
    try {
      const pool = await getSignerPool();
      const tx = await pool.Ban(workerKey);
      await tx.wait();
      alert('Worker banned');
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error banning worker');
    }
  };

  const unbanWorker = async () => {
    try {
      const pool = await getSignerPool();
      const tx = await pool.Unban(workerKey);
      await tx.wait();
      alert('Worker unbanned');
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error unbanning worker');
    }
  };

  const queryWorker = async () => {
    try {
      const pool = getReadPool();
      const wallet = await pool.GetWorkerAddress(workerKey);
      const approved: boolean = await pool.isApproved(workerKey);
      const banned: boolean = await pool.blacklist(workerKey);
      setWorkerInfo({ wallet, isApproved: approved, isBanned: banned });
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error querying worker');
    }
  };

  return (
    <div className="p-4 space-y-6">
      <h1 className="text-xl font-bold">Worker Management</h1>

      {isConnected ? (
        <span className="inline-block bg-green-500 text-white px-4 py-1 rounded">Connected</span>
      ) : (
        <button className="px-4 py-2 rounded shadow" onClick={connectWalletAndSwitchNetwork}>Connect & Switch Network</button>
      )}

      <section className="space-y-2 border p-4 rounded">
        <h3 className="font-semibold">Worker Key</h3>
        <input className="border p-2 rounded w-full" placeholder="Public key" value={workerKey} onChange={e => setWorkerKey(e.target.value)} />
        <div className="flex flex-wrap gap-2 pt-2">
          <button className="px-3 py-2 rounded shadow" onClick={registerWorker}>Register</button>
          <button className="px-3 py-2 rounded shadow" onClick={approveWorker}>Approve (owner)</button>
          <button className="px-3 py-2 rounded shadow" onClick={unregisterWorker}>Unregister</button>
          <button className="px-3 py-2 rounded shadow" onClick={banWorker}>Ban (owner)</button>
          <button className="px-3 py-2 rounded shadow" onClick={unbanWorker}>Unban (owner)</button>
          <button className="px-3 py-2 rounded shadow" onClick={queryWorker}>Query</button>
        </div>
      </section>

      <section className="space-y-2 border p-4 rounded">
        <h3 className="font-semibold">Update Worker Key</h3>
        <input className="border p-2 rounded w-full" placeholder="New public key" value={newWorkerKey} onChange={e => setNewWorkerKey(e.target.value)} />
        <button className="px-3 py-2 rounded shadow" onClick={updateWorker}>Update</button>
      </section>

      {workerInfo && (
        <section className="space-y-2 border p-4 rounded">
          <h3 className="font-semibold">Worker Info</h3>
          <pre className="bg-gray-100 p-2 rounded text-sm overflow-x-auto">
            {JSON.stringify(workerInfo, null, 2)}
          </pre>
        </section>
      )}
    </div>
  );
};

export default WorkerManager; 