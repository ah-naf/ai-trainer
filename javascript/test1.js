function findPairsWithSum(nums1, nums2, target) {
  const result = [];
  const set = new Set();

  for (const num of nums1) {
    set.add(num);
  }

  for (const num of nums2) {
    const complement = target - num;
    if (set.has(complement)) {
      result.push([complement, num]);
    }
  }

  return result;
}

let nums1 = [4, 5, 6, 7, 0, 1];
let nums2 = [3, 9, 10, 11, 12, 19];
let k = 13;

const pairs = findPairsWithSum(nums1, nums2, k);
console.log("Pairs that sum to", k, ":", pairs);
