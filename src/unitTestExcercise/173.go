package unitTestExcercise

type BSTIterator struct {
	record []*TreeNode
}

func Constructor1(root *TreeNode) BSTIterator {
	var iterator BSTIterator
	iterator.push(root)
	return iterator
}

func (this *BSTIterator) push(root *TreeNode) {
	for root != nil {
		this.record = append(this.record, root)
		root = root.Left
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	re := this.record[len(this.record)-1]
	this.record = this.record[:len(this.record)-1]
	this.push(re.Right)
	return re.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.record) != 0 {
		return true
	}
	return false
}
