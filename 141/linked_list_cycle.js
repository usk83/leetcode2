// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/
const main = () => {
  console.log("=========================================");
  console.log("No local tests available.");
  console.log("=========================================");
};

function ListNode(val) {
  this.val = val;
  this.next = null;
}

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
  let m = new Map();

  let objs = [
    {
      name: "aaa",
      age: 10,
    },
    {
      name: "bbb",
      age: 20,
    },
    {
      name: "ccc",
      age: 30,
    },
  ];

  for (let obj of objs) {
    m[obj] = true;
  }

  console.debug('=== DEBUG START ======================================');
  // m.forEach((value, key, map) => {
  //   console.debug(value, key, map);
  // });
  console.debug(m.size());
  console.debug('=== DEBUG END ========================================');

  return false;
};

main();
