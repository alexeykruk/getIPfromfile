## Task condition

    You need to parse a file with some lines that include the date, for example:

    01.01.2020 1.2.3.4 /index.html
    01.01.2020 1.2.3.1 /index.html
    01.01.2020 1.2.3.5 /index.html
    01.01.2021 1.2.3.4 /index.html
    01.01.2021 1.2.1.5 /index.html
    ...
    
    and should select the first 100 unique IP addresses

    The code includes:
        func uniqueIP(s []string) []string {
            ...
            // need for getting unique ip
        }

        func reverse(s []string) []string {
            ...
            // need for reverse the string
        }

        main block of code includes regexp condition for parse ip-address
        re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
