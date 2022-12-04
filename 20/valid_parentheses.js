// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/
const main = () => {
  // console.log("=========================================");
  // console.log(isValid("()"));
  // console.log(true);
  // console.log("=========================================");
  // console.log(isValid("()[]{}"));
  // console.log(true);
  // console.log("=========================================");
  // console.log(isValid("(]"));
  // console.log(false);
  console.log("=========================================");
  console.log(isValid("aa"));
  console.log(false);
  console.log("=========================================");
  // console.log(isValid("([)]"));
  // console.log(false);
  // console.log("=========================================");
  // console.log(isValid("{[]}"));
  // console.log(true);
  // console.log("=========================================");
};

var isValid = function(s) {
  let stack = [];
  let count = 0;
  for(let i = 0; i < s.length; i++) {
    console.log(i);
    switch (stack[stack.length - 1]) {
      case "(":
        if(s[i] === ")") {
          stack.pop();
        }
        break;
      case "{":
        if(s[i] === "}") {
          stack.pop();
        }
        break;
      case "[":
        if(s[i] === "]") {
          stack.pop();
        }
        break;
      case undefined:
        console.log("undefined");
      wow:
        console.log("wow!");
        // break wow;
        console.log("more!");

      default:
        console.log("push");
        stack.push(s[i]);
        break;
    }
    console.log(stack);
  }
  return !stack.length;
};

main();
