package _208_implement_trie_prefix_tree

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()

	// 测试插入apple后搜索apple应为true
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("Search('apple') returned false, want true")
	}

	// 测试插入apple后搜索app应为false
	if trie.Search("app") {
		t.Errorf("Search('app') returned true, want false")
	}

	// 测试插入apple后前缀匹配app应为true
	if !trie.StartsWith("app") {
		t.Errorf("StartsWith('app') returned false, want true")
	}

	// 测试插入app后搜索app应为true
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("Search('app') returned false after inserting 'app', want true")
	}

	// 测试插入空字符串
	trie.Insert("")
	if trie.Search("") {
		t.Errorf("Search('') returned true, want false (empty string not allowed)")
	}

	// 测试不存在的前缀
	if trie.StartsWith("xyz") {
		t.Errorf("StartsWith('xyz') returned true, want false")
	}

	// 测试部分匹配的前缀
	if !trie.StartsWith("ap") {
		t.Errorf("StartsWith('ap') returned false, want true")
	}

	// 测试完整匹配的前缀
	if !trie.StartsWith("apple") {
		t.Errorf("StartsWith('apple') returned false, want true")
	}

	// 测试超长前缀
	if trie.StartsWith("apples") {
		t.Errorf("StartsWith('apples') returned true, want false")
	}
}
