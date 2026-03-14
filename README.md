# golang-

LeetCode Hot 100 solutions in Go.

## Structure

Each problem is in its own package under `hot100/`:

```
hot100/p{number}_{name}/
  solution.go       # solution implementation
  solution_test.go  # test cases
```

## Running Tests

```bash
go test ./hot100/...
```

## Problems

| # | Problem | Category |
|---|---------|----------|
| 1 | [Two Sum](hot100/p0001_two_sum/) | Hash Map |
| 3 | [Longest Substring Without Repeating Characters](hot100/p0003_longest_substring_without_repeating_characters/) | Sliding Window |
| 10 | [Regular Expression Matching](hot100/p0010_regular_expression_matching/) | Dynamic Programming |
| 11 | [Container With Most Water](hot100/p0011_container_with_most_water/) | Two Pointers |
| 15 | [3Sum](hot100/p0015_3sum/) | Two Pointers |
| 17 | [Letter Combinations of a Phone Number](hot100/p0017_letter_combinations_of_a_phone_number/) | Backtracking |
| 19 | [Remove Nth Node From End of List](hot100/p0019_remove_nth_node_from_end_of_list/) | Linked List |
| 20 | [Valid Parentheses](hot100/p0020_valid_parentheses/) | Stack |
| 21 | [Merge Two Sorted Lists](hot100/p0021_merge_two_sorted_lists/) | Linked List |
| 22 | [Generate Parentheses](hot100/p0022_generate_parentheses/) | Backtracking |
| 23 | [Merge k Sorted Lists](hot100/p0023_merge_k_sorted_lists/) | Linked List / Heap |
| 31 | [Next Permutation](hot100/p0031_next_permutation/) | Array |
| 33 | [Search in Rotated Sorted Array](hot100/p0033_search_in_rotated_sorted_array/) | Binary Search |
| 34 | [Find First and Last Position of Element in Sorted Array](hot100/p0034_find_first_and_last_position_of_element_in_sorted_array/) | Binary Search |
| 39 | [Combination Sum](hot100/p0039_combination_sum/) | Backtracking |
| 42 | [Trapping Rain Water](hot100/p0042_trapping_rain_water/) | Two Pointers |
| 45 | [Jump Game II](hot100/p0045_jump_game_ii/) | Greedy |
| 46 | [Permutations](hot100/p0046_permutations/) | Backtracking |
| 48 | [Rotate Image](hot100/p0048_rotate_image/) | Array |
| 49 | [Group Anagrams](hot100/p0049_group_anagrams/) | Hash Map |
| 51 | [N-Queens](hot100/p0051_n_queens/) | Backtracking |
| 53 | [Maximum Subarray](hot100/p0053_maximum_subarray/) | Dynamic Programming |
| 54 | [Spiral Matrix](hot100/p0054_spiral_matrix/) | Array |
| 55 | [Jump Game](hot100/p0055_jump_game/) | Greedy |
| 56 | [Merge Intervals](hot100/p0056_merge_intervals/) | Array |
| 62 | [Unique Paths](hot100/p0062_unique_paths/) | Dynamic Programming |
| 64 | [Minimum Path Sum](hot100/p0064_minimum_path_sum/) | Dynamic Programming |
| 70 | [Climbing Stairs](hot100/p0070_climbing_stairs/) | Dynamic Programming |
| 72 | [Edit Distance](hot100/p0072_edit_distance/) | Dynamic Programming |
| 74 | [Search a 2D Matrix](hot100/p0074_search_a_2d_matrix/) | Binary Search |
| 75 | [Sort Colors](hot100/p0075_sort_colors/) | Two Pointers |
| 76 | [Minimum Window Substring](hot100/p0076_minimum_window_substring/) | Sliding Window |
| 78 | [Subsets](hot100/p0078_subsets/) | Backtracking |
| 79 | [Word Search](hot100/p0079_word_search/) | Backtracking |
| 84 | [Largest Rectangle in Histogram](hot100/p0084_largest_rectangle_in_histogram/) | Stack |
| 91 | [Decode Ways](hot100/p0091_decode_ways/) | Dynamic Programming |
| 94 | [Binary Tree Inorder Traversal](hot100/p0094_binary_tree_inorder_traversal/) | Tree |
| 98 | [Validate Binary Search Tree](hot100/p0098_validate_binary_search_tree/) | Tree |
| 100 | [Same Tree](hot100/p0100_same_tree/) | Tree |
| 101 | [Symmetric Tree](hot100/p0101_symmetric_tree/) | Tree |
| 102 | [Binary Tree Level Order Traversal](hot100/p0102_binary_tree_level_order_traversal/) | Tree / BFS |
| 104 | [Maximum Depth of Binary Tree](hot100/p0104_maximum_depth_of_binary_tree/) | Tree |
| 105 | [Construct Binary Tree from Preorder and Inorder Traversal](hot100/p0105_construct_binary_tree_from_preorder_and_inorder_traversal/) | Tree |
| 110 | [Balanced Binary Tree](hot100/p0110_balanced_binary_tree/) | Tree |
| 114 | [Flatten Binary Tree to Linked List](hot100/p0114_flatten_binary_tree_to_linked_list/) | Tree |
| 121 | [Best Time to Buy and Sell Stock](hot100/p0121_best_time_to_buy_and_sell_stock/) | Dynamic Programming |
| 124 | [Binary Tree Maximum Path Sum](hot100/p0124_binary_tree_maximum_path_sum/) | Tree |
| 128 | [Longest Consecutive Sequence](hot100/p0128_longest_consecutive_sequence/) | Hash Map |
| 131 | [Palindrome Partitioning](hot100/p0131_palindrome_partitioning/) | Backtracking |
| 136 | [Single Number](hot100/p0136_single_number/) | Bit Manipulation |
| 139 | [Word Break](hot100/p0139_word_break/) | Dynamic Programming |
| 141 | [Linked List Cycle](hot100/p0141_linked_list_cycle/) | Linked List |
| 142 | [Linked List Cycle II](hot100/p0142_linked_list_cycle_ii/) | Linked List |
| 143 | [Reorder List](hot100/p0143_reorder_list/) | Linked List |
| 146 | [LRU Cache](hot100/p0146_lru_cache/) | Design |
| 148 | [Sort List](hot100/p0148_sort_list/) | Linked List |
| 153 | [Find Minimum in Rotated Sorted Array](hot100/p0153_find_minimum_in_rotated_sorted_array/) | Binary Search |
| 155 | [Min Stack](hot100/p0155_min_stack/) | Stack |
| 160 | [Intersection of Two Linked Lists](hot100/p0160_intersection_of_two_linked_lists/) | Linked List |
| 169 | [Majority Element](hot100/p0169_majority_element/) | Array |
| 190 | [Reverse Bits](hot100/p0190_reverse_bits/) | Bit Manipulation |
| 191 | [Number of 1 Bits](hot100/p0191_number_of_1_bits/) | Bit Manipulation |
| 198 | [House Robber](hot100/p0198_house_robber/) | Dynamic Programming |
| 199 | [Binary Tree Right Side View](hot100/p0199_binary_tree_right_side_view/) | Tree / BFS |
| 200 | [Number of Islands](hot100/p0200_number_of_islands/) | Graph / BFS |
| 206 | [Reverse Linked List](hot100/p0206_reverse_linked_list/) | Linked List |
| 207 | [Course Schedule](hot100/p0207_course_schedule/) | Graph / Topological Sort |
| 208 | [Implement Trie (Prefix Tree)](hot100/p0208_implement_trie_prefix_tree/) | Trie |
| 210 | [Course Schedule II](hot100/p0210_course_schedule_ii/) | Graph / Topological Sort |
| 211 | [Design Add and Search Words Data Structure](hot100/p0211_design_add_and_search_words_data_structure/) | Trie |
| 212 | [Word Search II](hot100/p0212_word_search_ii/) | Trie / Backtracking |
| 213 | [House Robber II](hot100/p0213_house_robber_ii/) | Dynamic Programming |
| 215 | [Kth Largest Element in an Array](hot100/p0215_kth_largest_element_in_an_array/) | Heap |
| 221 | [Maximal Square](hot100/p0221_maximal_square/) | Dynamic Programming |
| 226 | [Invert Binary Tree](hot100/p0226_invert_binary_tree/) | Tree |
| 230 | [Kth Smallest Element in a BST](hot100/p0230_kth_smallest_element_in_a_bst/) | Tree |
| 234 | [Palindrome Linked List](hot100/p0234_palindrome_linked_list/) | Linked List |
| 235 | [Lowest Common Ancestor of a Binary Search Tree](hot100/p0235_lowest_common_ancestor_of_a_binary_search_tree/) | Tree |
| 238 | [Product of Array Except Self](hot100/p0238_product_of_array_except_self/) | Array |
| 239 | [Sliding Window Maximum](hot100/p0239_sliding_window_maximum/) | Sliding Window |
| 268 | [Missing Number](hot100/p0268_missing_number/) | Bit Manipulation |
| 283 | [Move Zeroes](hot100/p0283_move_zeroes/) | Two Pointers |
| 295 | [Find Median from Data Stream](hot100/p0295_find_median_from_data_stream/) | Heap |
| 297 | [Serialize and Deserialize Binary Tree](hot100/p0297_serialize_and_deserialize_binary_tree/) | Tree |
| 300 | [Longest Increasing Subsequence](hot100/p0300_longest_increasing_subsequence/) | Dynamic Programming |
| 309 | [Best Time to Buy and Sell Stock with Cooldown](hot100/p0309_best_time_to_buy_and_sell_stock_with_cooldown/) | Dynamic Programming |
| 312 | [Burst Balloons](hot100/p0312_burst_balloons/) | Dynamic Programming |
| 322 | [Coin Change](hot100/p0322_coin_change/) | Dynamic Programming |
| 323 | [Number of Connected Components](hot100/p0323_number_of_connected_components/) | Graph / Union Find |
| 338 | [Counting Bits](hot100/p0338_counting_bits/) | Bit Manipulation |
| 347 | [Top K Frequent Elements](hot100/p0347_top_k_frequent_elements/) | Heap |
| 394 | [Decode String](hot100/p0394_decode_string/) | Stack |
| 416 | [Partition Equal Subset Sum](hot100/p0416_partition_equal_subset_sum/) | Dynamic Programming |
| 417 | [Pacific Atlantic Water Flow](hot100/p0417_pacific_atlantic_water_flow/) | Graph / BFS |
| 437 | [Path Sum III](hot100/p0437_path_sum_iii/) | Tree |
| 438 | [Find All Anagrams in a String](hot100/p0438_find_all_anagrams_in_a_string/) | Sliding Window |
| 494 | [Target Sum](hot100/p0494_target_sum/) | Dynamic Programming |
| 518 | [Coin Change II](hot100/p0518_coin_change_ii/) | Dynamic Programming |
| 543 | [Diameter of Binary Tree](hot100/p0543_diameter_of_binary_tree/) | Tree |
| 560 | [Subarray Sum Equals K](hot100/p0560_subarray_sum_equals_k/) | Hash Map |
| 572 | [Subtree of Another Tree](hot100/p0572_subtree_of_another_tree/) | Tree |
| 684 | [Redundant Connection](hot100/p0684_redundant_connection/) | Graph / Union Find |
| 704 | [Binary Search](hot100/p0704_binary_search/) | Binary Search |
| 739 | [Daily Temperatures](hot100/p0739_daily_temperatures/) | Stack |
| 743 | [Network Delay Time](hot100/p0743_network_delay_time/) | Graph / Dijkstra |
| 763 | [Partition Labels](hot100/p0763_partition_labels/) | Greedy |
| 994 | [Rotting Oranges](hot100/p0994_rotting_oranges/) | Graph / BFS |
| 1143 | [Longest Common Subsequence](hot100/p1143_longest_common_subsequence/) | Dynamic Programming |
| 1448 | [Count Good Nodes in Binary Tree](hot100/p1448_count_good_nodes_in_binary_tree/) | Tree |