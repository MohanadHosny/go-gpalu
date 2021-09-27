# go-gpalu
**A simple [Gpalu](https://gpa.lu) wrapper made in Golang**

# Installation
```go get github.com/MohanadHosny/go-gpalu```

# Usage
## Generate new email
```golang
client := gpalu.NewClient(nil)
email := client.GetAddress()
```

## Get Inbox/Mails
```golang
client := gpalu.NewClient(nil)
mails := client.GetInbox("email@gpa.lu", 2)
```

# Notice
The **2** in GetInbox represents mails limit to grab. Also, it returns an **Array of struct Mail**.
<br/>

**Client.GetAddress** gives random email, tho you can use any random email **that includes @gpa.lu** without using this method.
