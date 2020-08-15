package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult 问题搜查结果返回结构体
type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue // Issue结构体切片
}

// Issue 问题结构体
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User 用户结构体
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues ：查询GitHub问题跟踪器
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	// 因为用户提供的查询条件可能包含类似?和&之类的特殊字符
	// 为了避免对URL造成冲突，我们用url.QueryEscape来对查询中的特殊字符进行转义操作
	q := url.QueryEscape(strings.Join(terms, " "))
	// 通过Get请求进行询问
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// 必须在所有执行路径上关闭resp.Body
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	// 使用了基于流式的解码器json.Decoder，它可以从一个输入流解码JSON数据
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
