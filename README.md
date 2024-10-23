# LeetList

## What is LeetList?

LeetList is a simple webapp service for creating and tracking your own list of LeetCode questions. This project is inspired by premade lists on the internet like [Grind 75](https://www.techinterviewhandbook.org/grind75) and [Neetcode 150](https://neetcode.io/practice). 

With LeetList, you can make your own customized lists and share them to everybody. You can make specific lists for different companies (based on what questions they've asked), different topics you want to get better at, different levels of LeetCode expertise... and so much more!

NOTE: LeetList is still a work in progress. The frontend interface is yet to be done :)

## Technologies Used

The backend API is sort of a funky mix of technologies. It's been implemented using:

- Golang
- GraphQL (via gqlgen)
- PostgreSQL
- Playwright (for webscraping)
  
The frontend is being implemented in:

- TypeScript
- React.js
- TailwindCSS
