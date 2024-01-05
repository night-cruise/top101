package recursionbacktracking

// 递归/回溯模板
// fn(状态空间，搜索位置，结束条件，访问数组，搜索路径，路径数组)
// 1. 判断结束条件，如果结束则保存搜索路径到路径数组中
// 2. 访问搜索位置，将搜索位置保存到访问数组中设置这个位置的搜索状态，将搜索位置的值保存到搜索路径中
// 3. 递归子节点，访问其他的搜索位置
// 4. 回溯，弹出搜索路径中保存的搜索位置的值，在访问数组中重置搜索位置的搜索状态
