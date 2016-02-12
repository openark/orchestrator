package discovery

// should write tests here

func z() {
	var (
		inSlice, outSlice []inst.InstanceKey
		ch                chan inst.InstanceKey
	)
	ch = make(chan inst.InstanceKey, 0)

	inSlice = []inst.InstanceKey{
		inst.InstanceKey{hostname: "a"}, // ignoring ports
		inst.InstanceKey{hostname: "b"},
		inst.InstanceKey{hostname: "c"},
		inst.InstanceKey{hostname: "d"},
		inst.InstanceKey{hostname: "e"},
	}

	mak

}
