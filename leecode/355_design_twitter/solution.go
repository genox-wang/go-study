package _55_design_twitter

/**
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户

*/

type Twitter struct {
	posts       []int
	idToPosts   map[int][]int
	idToFollows map[int]map[int]bool
}

// Constructor /** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		posts:       make([]int, 0),
		idToPosts:   make(map[int][]int, 0),
		idToFollows: make(map[int]map[int]bool, 0),
	}
}

// PostTweet /** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.posts = append(t.posts, tweetId)
	if _, ok := t.idToPosts[userId]; !ok {
		t.idToPosts[userId] = []int{}
	}
	t.idToPosts[userId] = append(t.idToPosts[userId], tweetId)
}

// GetNewsFeed /** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	postIds := map[int]bool{}
	if _, ok := t.idToPosts[userId]; ok {
		for _, id := range t.idToPosts[userId] {
			postIds[id] = true
		}
	}
	if _, ok := t.idToFollows[userId]; ok {
		for userId := range t.idToFollows[userId] {
			for _, id := range t.idToPosts[userId] {
				postIds[id] = true
			}
		}
	}
	ps := make([]int, 0)
	for i := len(t.posts) - 1; i >= 0; i-- {
		p := t.posts[i]
		if _, ok := postIds[p]; ok {
			ps = append(ps, p)
			if len(ps) >= 10 {
				break
			}
		}
	}
	return ps
}

// Follow /** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := t.idToFollows[followerId]; !ok {
		t.idToFollows[followerId] = map[int]bool{}
	}
	t.idToFollows[followerId][followeeId] = true
}

// Unfollow /** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := t.idToFollows[followerId]; ok {
		delete(t.idToFollows[followerId], followeeId)
	}
}
