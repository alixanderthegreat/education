I have used ruby before but I haven't haven't had a chance to work with it from the standpoint of having a pre-built test cases. 

I was curious on how I could call the `.rb` script without having to use `require` or `require_relative` in the test script. Not that this has particularly much value, at least not at the moment. Regardless, I thought that it would be neat to see if it were possible to require items for a script run from within the project directory without having to explicitly call them. Well, as I learned, you can use the following flags for the ruby interpreter:

```ruby
-rlibrary       require the library before executing your script
-Idirectory     specify $LOAD_PATH directory (may be used more than once)
```
And with that, you can craft a verticly descending script for the ruby interpreter to parse; hence, the following: 

```ruby
# (careful not to include trailing white space after forward-slash, eg `\ `)
ruby \
    -r minitest/autorun \
    -I . \
    -r hello_world \
    hello_world_test.rb
```
As a side note, regardless of whether `require`/`_relative` are explicitly in the test script, the test passes locally. As for the code passing with the [exercism](https://exercism.org) testing machine, it passes too. Furthermore, there really is no reason to do this, but it is interesting to consider ruby as a command line tool. 