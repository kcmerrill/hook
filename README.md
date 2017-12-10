[![Build Status](https://travis-ci.org/kcmerrill/hook.svg?branch=master)](https://travis-ci.org/kcmerrill/hook)
![hook](captain.hook.jpg)

# Hook

Add plugins, hooks, events or filter data within your go application.

```golang
Register("add-to-text", func(w *string) {
    *word += " More content added to the text."
})

Register("add-to-text", func(w *string) {
    *word += " This will be added too!"
})



words := "my simple text."

Filter("add-to-text", &words)

fmt.Println(words)
// my simple text. More content added to the text. This will be added too!
```

Want to add plugins? Instead of registering a function, register an executable command. `STDIN` will be passed as a json marshal'd interface that you pass in. The application will modify the contents to it's choosing, and then return. Assuming no errors, the new value will be unmarshal'd and set to the given interface.

```golang
Register("extra-text-to-word", "python plugin/python.py")
text := "hi"
Filter("extra-text-to-word", &text)
// hi-from-plugin
```