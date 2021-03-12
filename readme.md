# signaler

Tell your waiting goroutines that something happened in a threadsafe way.



### Backstory

I write a lot of code where I'll have one component that is waiting for some action to happen (i.e. user hits a button, data is received) and that needs to tell another component that it can check for this update.