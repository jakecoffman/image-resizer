[![Build Status](https://secure.travis-ci.org/jakecoffman/image-resizer.png?branch=master)](http://travis-ci.org/jakecoffman/image-resizer)


Created this for someone who needed to quickly resize many large images to 30% their original size.

1. Didn't use Python because I knew there was no way I could walk them through installing PIL on their machine
2. Wasn't sure if they had Java installed
3. Didn't want to use C or C++

Go was quick and easy as the std lib has image support and someone already implemented several resizing algorithms.
