# pwned

CLI tool based on the https://haveibeenpwned.com API to search for account or passwords included in data breaches

### Install
`go get github.com/mvording/pwned` <br>
If Go environment is default, navigate to ~/go/src/github.com/mvording/pwned and: <br>
`go install pwned.go`

### Usage

#### For one request:
`pwned` specifying either account or password parameter 

##### Breach Search
when specifying the `-account` flag (usually email address) it will return the breaches that included the given account 

##### Password Search
when specifying the `-password` flag, it will return the number of times the password was found online in data breaches


#### For batch requests

You can specify a set of accounts or passwords by specifying the `-mode` flag.
`-mode=breach` breach search
`-mode=password` password count search
To specify a payload you can either redirect/pipe, or explicitly provide a filename with `-file="path/to/file.txt"`


#### Format
By default the output is in text, you can specify json format for either mode with `-format=json`. For breach requests this will output the full set of breach properties.


#### Examples 

```
./pwned -account=hello@microsoft.com 
hello@microsoft.com=11 // Adobe Apollo Disqus Evony LinkedIn MySpace Neopets Tianya TrikSpamBotnet VerificationsIO Youku

./pwned -password=passw0rd
passw0rd=216221



./pwned -mode=password -file="./test.pass"
passw0rd=216221
password=3645804

./pwned -mode=breach -file="./test.breach"
hello@microsoft.com=11 // Adobe Apollo Disqus Evony LinkedIn MySpace Neopets Tianya TrikSpamBotnet VerificationsIO Youku 
hello@google.com=26 // Apollo CashCrate Dailymotion Dubsmash EverybodyEdits Evony FlashFlashRevolution HauteLook HeroesOfNewerth Houzz iDressup iMesh Lastfm Leet Lifeboat MyFitnessPal MyHeritage MySpace NextGenUpdate OnlinerSpambot RiverCityMedia VerificationsIO Youku Pastebin(g35Jndue) Pastebin(E4qKaR2B) Pastebin(TNpbvQpR) 

./pwned -mode=breach < test.breach
same as above

cat test.breach | ./pwned -mode=breach 
same as above
```
