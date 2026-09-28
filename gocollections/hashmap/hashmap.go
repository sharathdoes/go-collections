package hashmap

type Entry struct {
	Key string
	Value string
}

type HashMap struct {
	items [][]Entry
}

func NewHashMap() *HashMap {
	buckets := make([][]Entry, 10)
	return &HashMap{ items:buckets}
}

func(h *HashMap) Put(key, value string) {

	hashed:=hash(key)
	bucket:=hashed % 10
	if h.Get(key)==""{
		h.items[bucket]=append(h.items[bucket], Entry{key, value})
	} else {
		h.Update(bucket, key,value)
	}
}

func(h *HashMap) Get(key string) string {
	bucket:= h.items[hash(key)%10]
	for _, k:=range bucket {
		if key==k.Key {
			return k.Value
		}
	}
	return ""
}

func(h *HashMap) Update(bucket int, key string, value string)  {
	
	// for _, k:=range h.items[bucket] {
	// 	if key==k.Key {
	// 		k.Value=value 
	// 		return 
	// 	}
	// } doesnt work because k is a copy of slice element
	entries:=h.items[bucket]
	for i := range entries {
		if entries[i].Key == key {
			entries[i].Value = value
		}
	}
}

func hash(input string ) int {
	hash := 0
	for _, char := range input {
		hash = hash*31 + int(char)
	}
	return hash
}