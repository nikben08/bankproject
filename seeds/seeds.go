package seeds

func SuperUser() map[string]string {
	var user = map[string]string{
		"username":    "admin",
		"hash":        "08f46d939d7ff7ebb1df9de1cc246135c7a8694cabf0e1c1cc67c50e12832f08",
		"accessLevel": "1",
	}
	return user
}

func Credits() map[int]string {
	var credits = map[int]string{
		1: "Konut",
		2: "Tuketici",
		3: "Mevduat",
	}
	return credits
}

func TimeOptions() map[int]string {
	var timeOptions = map[int]string{
		1: "3 ay",
		2: "6 ay",
		3: "12 ay",
		4: "24 ay",
		5: "36 ay",
		6: "5 yıl",
		7: "10 yıl",
	}
	return timeOptions
}
