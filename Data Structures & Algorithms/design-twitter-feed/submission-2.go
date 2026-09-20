type Twitter struct {
	following map[int]map[int]struct{}
	tweets map[int][]Tweet
	time int 
    
}
type Tweet struct{
	time int
	tweetId int 
}
type HeapItem struct{
	tweet Tweet
	userId int 
	listIndex int 
}


func Constructor() Twitter {
	following := make(map[int]map[int]struct{})
	tweets := make(map[int][]Tweet)
	time := 0 
	return Twitter {
		following : following,
		tweets: tweets,
		time: time,

	}  
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.time++
	newTweet := Tweet{
		time: this.time,
		tweetId: tweetId,
	}
	this.tweets[userId] = append(this.tweets[userId], newTweet)
    
}


func (this *Twitter) GetNewsFeed(userId int) []int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	retSlice := []int{}

	ownTweets := this.tweets[userId]

	if len(ownTweets) > 0 {
    	lastIndex := len(ownTweets) - 1

    	heap.Push(maxHeap, HeapItem{
        	tweet:     ownTweets[lastIndex],
        	userId:    userId,
        	listIndex: lastIndex,
    	})
	}

	for followeeId := range this.following[userId]{
		followeeTweets := this.tweets[followeeId]
		if len(followeeTweets) > 0{
			lastIndex := len(followeeTweets)-1

			item := HeapItem{
				tweet: followeeTweets[lastIndex],
				userId: followeeId,
				listIndex: lastIndex,
			}
			heap.Push(maxHeap,item)
		}

	}
	for i := 0; i < 10 && maxHeap.Len() > 0; i++{
		item := heap.Pop(maxHeap).(HeapItem)
		retSlice = append(retSlice,item.tweet.tweetId)

		if item.listIndex > 0{
			nextIndex := item.listIndex -1 

			newItem := HeapItem{
				tweet: this.tweets[item.userId][nextIndex],
				userId: item.userId,
				listIndex: nextIndex,
			}
			heap.Push(maxHeap,newItem)
		}
	}
return retSlice 
    
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
	if this.following[followerId] == nil {
        this.following[followerId] = make(map[int]struct{})
    }
    this.following[followerId][followeeId] = struct{}{}
}
    



func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	delete(this.following[followerId],followeeId)
}

type MaxHeap []HeapItem

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { 
    return h[i].tweet.time > h[j].tweet.time 
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
   *h = append(*h, x.(HeapItem))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

