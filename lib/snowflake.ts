const EPOCH = 1735689601; // Jan 1st, 2025 00:00:01 GMT+0000
const MACHINE_ID_BITS = 10; // machine id bits
const SEQUENCE_BITS = 12; // sequence bits
const MACHINE_ID_SHIFT = SEQUENCE_BITS;
const TIMESTAMP_SHIFT = SEQUENCE_BITS + MACHINE_ID_BITS;
const MAX_MACHINE_ID = -1 ^ (-1 << MACHINE_ID_BITS);
const MAX_SEQUENCE = -1 ^ (-1 << SEQUENCE_BITS);

class Snowflake {
  private static instance: Snowflake;

  static getInstance() {
    if (!this.instance) {
      this.instance = new Snowflake();
    }

    return this.instance;
  }

  private machineId = 0;
  private lastTime = 0;
  private sequence = 0;

  generate(): number {
    let now = new Date().getTime();

    if (now === this.lastTime) {
      this.sequence = (this.sequence + 1) & MAX_SEQUENCE;

      // handle sequence overflow
      if (this.sequence === 0) {
        // block until the next millisecond
        while (now <= this.lastTime) {
          now = new Date().getTime();
        }
      }
    } else {
      this.sequence = 0;
    }

    this.lastTime = now;

    return Number(
      (BigInt(now - EPOCH) << BigInt(TIMESTAMP_SHIFT)) |
        (BigInt(this.machineId) << BigInt(MACHINE_ID_SHIFT)) |
        BigInt(this.sequence)
    );
  }
}

export const snowflake = Snowflake.getInstance();
