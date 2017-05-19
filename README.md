# goChat
> ### Simple TCP chat implementation in Go.

## Usage: 

### 1. Start Server
##### To start, use : `go run ./server.go`
###### *By default it runs at port 6000*

___

### 2. Connect to server via telnet
#### a) Same device : `telnet localhost 6000`
#### b) Different device on same network : `telnet [local IP] 6000`
#### c) Externally : `telnet [public IP] 6000`
###### *(For external connections you need to have port forwarding on)*

___

### 3. Chat

> #### *In this conversation Kalbhor connects first and Andrew disconnects first.*

#### Kalbhor's chat box

```
❯
Escape character is '^]'.
Enter name: kalbhor
Accepted user : [kalbhor]

Accepted user : [Andrew Ng]

Hey Andrew!
>kalbhor: Hey Andrew!

>Andrew Ng: Hey, kalbhor. What's up?

I'm planning on taking your course on Machine learning!
>kalbhor: I'm planning on taking your course on Machine learning!

>Andrew Ng: That's great! I'll go now. Busy teaching machines.

Andrew Ng disconnected
```

#### Andrew's chat box

```
❯
Escape character is '^]'.
Enter name: Andrew Ng
Accepted user : [Andrew Ng]

>kalbhor: Hey Andrew!

Hey, kalbhor. What's up?
>Andrew Ng: Hey, kalbhor. What's up?

>kalbhor: I'm planning on taking your course on Machine learning!

That's great! I'll go now. Busy teaching machines.
>Andrew Ng: That's great! I'll go now. Busy teaching machines.
```

___

## License

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/93/GPLv3_Logo.svg/1200px-GPLv3_Logo.svg.png" alt="license" height="40%" width="40%">


