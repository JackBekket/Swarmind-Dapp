const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Pool", function () {
  let owner, user1, other;
  let pool, nft, credit;

  beforeEach(async function () {
    [owner, user1, other] = await ethers.getSigners();

    const LLMNFT = await ethers.getContractFactory("LLMNFT");
    nft = await LLMNFT.deploy("LLM Models", "LLM");

 
    const Cred = await ethers.getContractFactory("Cred");
    credit = await Cred.deploy(1000000000000000);


    const FeesCalculator = await ethers.getContractFactory("FeesCalculator");
    const feesCalculator = await FeesCalculator.deploy();

    const Pool = await ethers.getContractFactory("Pool", {
      libraries: {
        "contracts/FeesCalculator.sol:FeesCalculator": feesCalculator.target,
      },
    });
    pool = await Pool.deploy(nft.target, credit.target);
  });

  it("should register a new worker", async function () {
    const laiKey = "test-lai-key";

    await pool.connect(user1).RegisterWorker(laiKey);

    const workerAddr = await pool.GetWorkerAddress(laiKey);
    expect(workerAddr).to.equal(user1.address);

    const approved = await pool.isApproved(laiKey);
    expect(approved).to.be.false;

    await expect(pool.connect(user1).RegisterWorker("another-key"))
      .to.emit(pool, "NewWorker")
      .withArgs(user1.address, false);
  });

  it("should reject already registered key", async function () {
    const laiKey = "test-lai-key";

    await pool.connect(user1).RegisterWorker(laiKey);
    await expect(pool.connect(user1).RegisterWorker(laiKey)).to.be.revertedWith(
      "Key already registered"
    );
  });

  it("should reject blacklisted keys", async function () {
    const badKey = "blacklist-me";

    await pool.connect(owner).Ban(badKey);
    await expect(pool.connect(user1).RegisterWorker(badKey)).to.be.revertedWith(
      "This key is in blacklist"
    );
  });

  it("should not allow more than 20 workers per wallet", async function () {
    for (let i = 0; i < 20; i++) {
      const key = `key${i}`;
      await pool.connect(user1).RegisterWorker(key);
    }

    await expect(
      pool.connect(user1).RegisterWorker("key20")
    ).to.be.revertedWith("Maximum number of workers (20) reached");
  });
});

