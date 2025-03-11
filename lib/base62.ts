const BASE62_CHARSET =
  "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

export function tob62(id: number): string {
  if (id < 62) {
    return BASE62_CHARSET[id];
  }

  const b = [];
  while (id >= 62) {
    b.push(BASE62_CHARSET[id % 62]);
    id = Math.floor(id / 62);
  }

  return b.reverse().join("");
}
