# msg – a simple utility for sending sms

This is a very simple command line utility that just sends an SMS to a
given number using Twilio.  It doesn't try to do anything fancy.

# Pre-built releases

Check out https://github.com/borud/msg/releases

# Installing

In order to build this utility you need Kevin Burke's Twilio library:

    go get github.com/kevinburke/twilio-go
	
Then you just build and install it.

    go install
	
Before using it you have to add the access credentials to your
environment.  You can do this by adding the variables to your default
login environment or by setting them in the shellscript if you use
this utility in a shellscript.

    export TWILIO_SID="<your sid>"
	export TWILIO_TOKEN="<your token>"
	
# Running

If you have the `TWILIO_SID` and `TWILIO_TOKEN` in your environment
you can just run the utility like this:

    msg <phone number> "<message>"
	
And that's it.


