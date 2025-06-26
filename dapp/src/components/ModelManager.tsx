import React, { useState } from 'react';
import { ethers } from 'ethers';
import {
  llmnftAddress,
  poolAddress,
  chainId,
  chainName,
  curName,
  curSymbol,
  curDec,
  chainRPC,
  blockExplorer,
} from '../constants';
import { llmnftAbi, poolAbi } from '../abis';

interface ModelData {
  tokenId: string;
  authorRoyalty: string;
  authorWallet: string;
  hwPriceIn: string;
  hwPriceOut: string;
  payType: string;
}

const ModelManager: React.FC = () => {
  // Connection state
  const [isConnected, setIsConnected] = useState(false);

  // Create NFT inputs
  const [hfId, setHfId] = useState('');
  const [royaltyPrice, setRoyaltyPrice] = useState('');
  const [modelType, setModelType] = useState<'text' | 'image'>('text');
  const [maxContext, setMaxContext] = useState('');

  // Add model inputs
  const [tokenIdForAdd, setTokenIdForAdd] = useState('');
  const [hwPriceIn, setHwPriceIn] = useState('');
  const [hwPriceOut, setHwPriceOut] = useState('');

  // Get / Update / Delete inputs
  const [modelIdQuery, setModelIdQuery] = useState('');
  const [fetchedModel, setFetchedModel] = useState<ModelData | null>(null);

  const [newRoyalty, setNewRoyalty] = useState('');
  const [newAuthorAddress, setNewAuthorAddress] = useState('');

  const provider = typeof window !== 'undefined' && (window as any).ethereum ? new ethers.BrowserProvider((window as any).ethereum) : null;

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
    } catch (error) {
      console.error(error);
      alert('Failed to connect or switch network.');
    }
  };

  // Helpers
  const getSignerContracts = async () => {
    if (!provider) throw new Error('No provider');
    const signer = await provider.getSigner();
    const nft = new ethers.Contract(llmnftAddress, llmnftAbi, signer);
    const pool = new ethers.Contract(poolAddress, poolAbi, signer);
    return { signer, nft, pool };
  };

  const handleCreateNFT = async () => {
    try {
      const { nft, signer } = await getSignerContracts();
      const wallet = await signer.getAddress();
      const price = BigInt(royaltyPrice);
      const mType = modelType === 'text' ? 0 : 1; // enum mapping
      const maxCtx = BigInt(maxContext);

      const tx = await nft.CreateLLM_NFT(hfId, price, wallet, mType, maxCtx);
      const receipt = await tx.wait();
      const tokenId = receipt?.logs?.[0]?.topics?.[3] || 'unknown';
      alert(`NFT created! Token ID: ${tokenId}`);
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error creating NFT');
    }
  };

  const handleAddModel = async () => {
    try {
      const { pool } = await getSignerContracts();
      const tx = await pool.AddModel(BigInt(tokenIdForAdd), BigInt(hwPriceIn), BigInt(hwPriceOut));
      await tx.wait();
      alert('Model added to pool!');
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error adding model');
    }
  };

  const handleGetModel = async () => {
    try {
      if (!provider) throw new Error('No provider');
      const pool = new ethers.Contract(poolAddress, poolAbi, provider);
      const data = await pool.GetModel(BigInt(modelIdQuery));
      const formatted: ModelData = {
        tokenId: data.token_id.toString(),
        authorRoyalty: data.author_royalty.toString(),
        authorWallet: data.author_wallet,
        hwPriceIn: data.hw_price_per_input_token.toString(),
        hwPriceOut: data.hw_price_per_output_token.toString(),
        payType: data.pay_type_.toString(),
      };
      setFetchedModel(formatted);
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error fetching model');
    }
  };

  const handleUpdateModel = async () => {
    try {
      const { pool } = await getSignerContracts();
      const tx = await pool.UpdateModel(BigInt(modelIdQuery), BigInt(newRoyalty), newAuthorAddress);
      await tx.wait();
      alert('Model updated');
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error updating model');
    }
  };

  const handleDeleteModel = async () => {
    try {
      const { pool } = await getSignerContracts();
      const tx = await pool.DeleteModel(BigInt(modelIdQuery));
      await tx.wait();
      alert('Model deleted');
    } catch (err: any) {
      console.error(err);
      alert(err?.message || 'Error deleting model');
    }
  };

  return (
    <div className="p-4 space-y-6">
      <h1 className="text-xl font-bold">Model Management</h1>

      {isConnected ? (
        <span className="inline-block bg-green-500 text-white px-4 py-1 rounded">Connected</span>
      ) : (
        <button className="px-4 py-2 rounded shadow" onClick={connectWalletAndSwitchNetwork}>Connect & Switch Network</button>
      )}

      {/* Create NFT */}
      <section className="space-y-2 border p-4 rounded">
        <h3 className="font-semibold">Create LLM NFT</h3>
        <input className="border p-2 rounded w-full" placeholder="HuggingFace ID" value={hfId} onChange={e => setHfId(e.target.value)} />
        <input className="border p-2 rounded w-full" placeholder="Royalty price" type="number" value={royaltyPrice} onChange={e => setRoyaltyPrice(e.target.value)} />
        <div className="flex items-center space-x-2">
          <label>Model Type:</label>
          <select value={modelType} onChange={e => setModelType(e.target.value as 'text' | 'image')} className="border p-2 rounded">
            <option value="text">Text</option>
            <option value="image">Image</option>
          </select>
        </div>
        <input className="border p-2 rounded w-full" placeholder="Max Context Window" type="number" value={maxContext} onChange={e => setMaxContext(e.target.value)} />
        <button onClick={handleCreateNFT} className="px-4 py-2 rounded shadow">Create NFT</button>
      </section>

      {/* Add Model to Pool */}
      <section className="space-y-2 border p-4 rounded">
        <h3 className="font-semibold">Add Model to Pool</h3>
        <input className="border p-2 rounded w-full" placeholder="Token ID" type="number" value={tokenIdForAdd} onChange={e => setTokenIdForAdd(e.target.value)} />
        <input className="border p-2 rounded w-full" placeholder="HW Price In" type="number" value={hwPriceIn} onChange={e => setHwPriceIn(e.target.value)} />
        <input className="border p-2 rounded w-full" placeholder="HW Price Out" type="number" value={hwPriceOut} onChange={e => setHwPriceOut(e.target.value)} />
        <button onClick={handleAddModel} className="px-4 py-2 rounded shadow">Add Model</button>
      </section>

      {/* Get / Update / Delete Model */}
      <section className="space-y-2 border p-4 rounded">
        <h3 className="font-semibold">Manage Existing Model</h3>
        <input className="border p-2 rounded w-full" placeholder="Model ID" type="number" value={modelIdQuery} onChange={e => setModelIdQuery(e.target.value)} />
        <div className="flex space-x-2">
          <button onClick={handleGetModel} className="px-4 py-2 rounded shadow">Get</button>
          <button onClick={handleUpdateModel} className="px-4 py-2 rounded shadow">Update</button>
          <button onClick={handleDeleteModel} className="px-4 py-2 rounded shadow">Delete</button>
        </div>
        {fetchedModel && (
          <pre className="bg-gray-100 p-2 rounded text-sm overflow-x-auto">{JSON.stringify(fetchedModel, null, 2)}</pre>
        )}
        {/* Update fields */}
        <div className="pt-2 space-y-2">
          <input className="border p-2 rounded w-full" placeholder="New Royalty" type="number" value={newRoyalty} onChange={e => setNewRoyalty(e.target.value)} />
          <input className="border p-2 rounded w-full" placeholder="New Author Wallet" value={newAuthorAddress} onChange={e => setNewAuthorAddress(e.target.value)} />
        </div>
      </section>
    </div>
  );
};

export default ModelManager; 