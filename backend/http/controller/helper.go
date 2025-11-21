package controller

func BuildPolicyJSON(role string) []byte {
	switch role {
	case "admin":
		return []byte(`{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Action": ["s3:*"],
			"Resource": []
		}
	]
}`)
	case "user":
		return []byte(`{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Action": [
				"s3:GetObject",
				"s3:PutObject",
				"s3:DeleteObject",
				"s3:ListBucket"
			],
			"Resource": []
		}
	]
}`)
	case "viewer":
		return []byte(`{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Action": ["s3:GetObject", "s3:ListBucket"],
			"Resource": []
		}
	]
}`)
	default:
		return []byte(`{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Action": ["s3:*"],
			"Resource": []
		}
	]
}`)
	}
}
