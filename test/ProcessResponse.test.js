const { expect, assert } = require("chai");
const { ethers, network } = require("hardhat");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const fs = require("fs");

/**
 * Test suite for Pool.ProcessResponse covering:
 *  - Positive flows for token-based & request-based pricing
 *  - Access control & validation reverts
 *  - Edge-cases (zero tokens, large numbers)
 */
describe("Pool – ProcessResponse", function () {
  let owner;
  let worker; // will own the worker key
  let author; // will receive author royalty
  let stranger; // random account for negative tests

  let pool;
  let nft;
  let credit;
  let feesCalculator;

  const initialCreditSupply = ethers.parseUnits("1000000", 9); // 1M SWM (decimals = 9)

  const gasStats = {};
  function recordGas(label, gas) {
    if (!gasStats[label]) gasStats[label] = [];
    gasStats[label].push(gas.toString());
  }

  // helpful helper to deploy whole stack once for each test case
  async function deployFixture() {
    [owner, worker, author, stranger] = await ethers.getSigners();

    // Deploy LLM NFT collection
    const LLMNFT = await ethers.getContractFactory("LLMNFT");
    nft = await LLMNFT.deploy("LLM Models", "LLM");

    // Deploy ERC20 token used as credit
    const Cred = await ethers.getContractFactory("Cred");
    credit = await Cred.deploy(initialCreditSupply);

    // Mint additional tokens for owner to fund Pool later (constructor already minted to owner)

    // Deploy FeesCalculator library
    const FeesCalculator = await ethers.getContractFactory("FeesCalculator");
    feesCalculator = await FeesCalculator.deploy();

    // Deploy Pool with linked library
    const Pool = await ethers.getContractFactory("Pool", {
      libraries: {
        "contracts/FeesCalculator.sol:FeesCalculator": feesCalculator.target,
      },
    });
    pool = await Pool.deploy(nft.target, credit.target);
    await pool.waitForDeployment();

    return { pool, nft, credit };
  }

  // Convenience method to set up an LLM model & add to Pool metadata
  async function createModel({ modelType, royaltyPerToken, hwPriceIn, hwPriceOut }) {
    // create LLM NFT (token id = 1,2,...)
    const tx = await nft.CreateLLM_NFT(
      "model/hf/id",
      royaltyPerToken,
      author.address,
      modelType, // 0=text,1=image
      2048 // max context window (unused here)
    );
    const recMint = await tx.wait();
    recordGas("CreateLLM_NFT", recMint.gasUsed);
    const tokenId = 1n; // first mint in fixture

    // Add model to Pool (onlyOwner)
    const txAdd = await pool.AddModel(tokenId, hwPriceIn, hwPriceOut);
    const recAdd = await txAdd.wait();
    recordGas("AddModel", recAdd.gasUsed);
    return tokenId;
  }

  // Registers and approves a worker (helper)
  async function registerAndApproveWorker(laiKey) {
    const txReg = await pool.connect(worker).RegisterWorker(laiKey);
    const recReg = await txReg.wait();
    recordGas("RegisterWorker", recReg.gasUsed);
    const txApp = await pool.ApproveWorker(laiKey);
    const recApp = await txApp.wait();
    recordGas("ApproveWorker", recApp.gasUsed);
  }

  beforeEach(async function () {
    await deployFixture();
  });

  describe("Positive flows", function () {
    it("should process response for token-based pricing and transfer correct funds", async function () {
      // --- Arrange
      const laiKey = "worker-key-token";
      await registerAndApproveWorker(laiKey);

      // create TEXT model (pay_type = token)
      const royalty = 3n;
      const hwpIn = 2n;
      const hwpOut = 4n;
      const tokenId = await createModel({ modelType: 0, royaltyPerToken: royalty, hwPriceIn: hwpIn, hwPriceOut: hwpOut });

      // parameters for call
      const inputT = 10n;
      const outputT = 5n;
      const tokensSum = inputT + outputT;
      const costHW = hwpIn * inputT + hwpOut * outputT; // 2*10 + 4*5 = 40
      const authorCost = royalty * tokensSum; // 3*15 = 45
      const totalCost = costHW + authorCost; // 85

      // fund Pool with enough tokens
      await credit.transfer(await pool.getAddress(), totalCost);

      const workerBalBefore = await credit.balanceOf(worker.address);
      const authorBalBefore = await credit.balanceOf(author.address);
      const poolBalBefore = await credit.balanceOf(await pool.getAddress());

      // --- Act & Assert events
      const txPR = await pool.ProcessResponse(1, laiKey, tokenId, inputT, outputT, 100);
      const recPR = await txPR.wait();
      recordGas("ProcessResponse_token", recPR.gasUsed);
      await expect(txPR)
        .to.emit(pool, "Response")
        .withArgs(
          1,
          tokenId,
          worker.address,
          inputT,
          outputT,
          totalCost,
          100,
          anyValue // timestamp
        )
        .and.to.emit(pool, "Payout")
        .withArgs(worker.address, tokenId, costHW);

      // --- Assert balances
      const workerBalAfter = await credit.balanceOf(worker.address);
      const authorBalAfter = await credit.balanceOf(author.address);
      const poolBalAfter = await credit.balanceOf(await pool.getAddress());

      expect(workerBalAfter - workerBalBefore).to.equal(costHW);
      expect(authorBalAfter - authorBalBefore).to.equal(authorCost);
      expect(poolBalBefore - poolBalAfter).to.equal(totalCost);
    });

    it("should process response for request-based pricing and transfer correct funds", async function () {
      const laiKey = "worker-key-request";
      await registerAndApproveWorker(laiKey);

      // IMAGE model (pay_type = request)
      const royalty = 20n;
      const hwPerRequest = 100n; // passed as hwPriceIn
      const tokenId = await createModel({ modelType: 1, royaltyPerToken: royalty, hwPriceIn: hwPerRequest, hwPriceOut: 0n });

      const costHW = hwPerRequest;
      const authorCost = royalty;
      const totalCost = costHW + authorCost; // 120

      await credit.transfer(await pool.getAddress(), totalCost);

      const workerBalBefore = await credit.balanceOf(worker.address);
      const authorBalBefore = await credit.balanceOf(author.address);

      const txPR2 = await pool.ProcessResponse(42, laiKey, tokenId, 0, 0, 77);
      const recPR2 = await txPR2.wait();
      recordGas("ProcessResponse_request", recPR2.gasUsed);
      await expect(txPR2)
        .to.emit(pool, "Response")
        .withArgs(42, tokenId, worker.address, 0, 0, totalCost, 77, anyValue)
        .and.to.emit(pool, "Payout")
        .withArgs(worker.address, tokenId, costHW);

      const workerBalAfter = await credit.balanceOf(worker.address);
      const authorBalAfter = await credit.balanceOf(author.address);

      expect(workerBalAfter - workerBalBefore).to.equal(costHW);
      expect(authorBalAfter - authorBalBefore).to.equal(authorCost);
    });
  });

  describe("Negative cases", function () {
    it("should revert if worker is not approved", async function () {
      const laiKey = "not-approved-worker";
      await pool.connect(worker).RegisterWorker(laiKey);
      // not calling ApproveWorker

      const tokenId = await createModel({ modelType: 0, royaltyPerToken: 1n, hwPriceIn: 1n, hwPriceOut: 1n });
      await expect(
        pool.ProcessResponse(1, laiKey, tokenId, 1, 1, 1)
      ).to.be.revertedWith("This worker is not yet approved");
    });

    it("should revert if worker is blacklisted", async function () {
      const laiKey = "blacklisted-worker";
      await registerAndApproveWorker(laiKey);
      await pool.Ban(laiKey);

      const tokenId = await createModel({ modelType: 0, royaltyPerToken: 1n, hwPriceIn: 1n, hwPriceOut: 1n });
      await expect(
        pool.ProcessResponse(1, laiKey, tokenId, 1, 1, 1)
      ).to.be.revertedWith("This worker is in blacklist");
    });

    it("should revert for unregistered worker id", async function () {
      const fakeKey = "ghost-worker";
      const tokenId = await createModel({ modelType: 0, royaltyPerToken: 1n, hwPriceIn: 1n, hwPriceOut: 1n });
      await expect(
        pool.ProcessResponse(1, fakeKey, tokenId, 1, 1, 1)
      ).to.be.revertedWith("This worker is not yet approved");
    });
  });

  describe("Edge cases", function () {
    it("should work with zero input/output tokens (zero cost)", async function () {
      const laiKey = "zero-tokens";
      await registerAndApproveWorker(laiKey);

      const royalty = 5n;
      const hwpIn = 2n;
      const hwpOut = 3n;
      const tokenId = await createModel({ modelType: 0, royaltyPerToken: royalty, hwPriceIn: hwpIn, hwPriceOut: hwpOut });

      // No need to fund pool – cost must be 0

      const txZero = await pool.ProcessResponse(99, laiKey, tokenId, 0, 0, 15);
      const recZero = await txZero.wait();
      recordGas("ProcessResponse_zero", recZero.gasUsed);
      await expect(txZero)
        .to.emit(pool, "Response")
        .withArgs(99, tokenId, worker.address, 0, 0, 0, 15, anyValue)
        .and.to.emit(pool, "Payout")
        .withArgs(worker.address, tokenId, 0);
    });

    it("should handle large numbers without overflow", async function () {
      const laiKey = "big-numbers";
      await registerAndApproveWorker(laiKey);

      const royalty = 10n;
      const hwpIn = 20n;
      const hwpOut = 30n;
      const tokenId = await createModel({ modelType: 0, royaltyPerToken: royalty, hwPriceIn: hwpIn, hwPriceOut: hwpOut });

      const input = 1_000_000n;
      const output = 2_000_000n;
      const costHW = hwpIn * input + hwpOut * output;
      const authorCost = royalty * (input + output);
      const totalCost = costHW + authorCost;

      await credit.transfer(await pool.getAddress(), totalCost);

      const txBig = await pool.ProcessResponse(123, laiKey, tokenId, input, output, 500);
      const recBig = await txBig.wait();
      recordGas("ProcessResponse_big", recBig.gasUsed);
      await expect(txBig)
        .to.emit(pool, "Response")
        .withArgs(123, tokenId, worker.address, input, output, totalCost, 500, anyValue);
    });
  });

  after(async function () {
    const file = "gas-stats.json";
    let data = {};
    if (fs.existsSync(file)) {
      try {
        data = JSON.parse(fs.readFileSync(file));
      } catch (e) {
        // ignore parse errors, overwrite
      }
    }
    Object.entries(gasStats).forEach(([k, v]) => {
      if (data[k]) {
        data[k] = data[k].concat(v);
      } else {
        data[k] = v;
      }
    });
    fs.writeFileSync(file, JSON.stringify(data, null, 2));
  });
}); 