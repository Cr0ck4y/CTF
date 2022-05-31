# same old

1. Firstly I've tried to read about CRC32 details. I've seen that everyone is telling about that it's insecure. So decided to slightly change my tactic. Let's go with brutforcing.
2. I've gathered all available charakters `{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}`
3. We firstly need to put our team name and then add some alphanumeric characters at the end of it so that the crc32 output of that string is same as crc32 output of word "the" (missing in `Hack the planet!`)

I've tried to write simple script with python [`solution`](same-old/solution.py) and later rewrite it in golang just bc it is fun [`solution`](same-old/solution.go) :)))

After submitting any of prepared strings we get correct submition 🖤
