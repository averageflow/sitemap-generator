# Sitemap Generator

This CLI tool was born out of the need to create an effective crawler that would generate a sitemap for websites.

Usage:


```sh
# Normal mode (-multilingual=false)
./sitemap-generator -location=https://www.mywebsite.com 
```

```bash
# Multilingual mode (-multilingual=true is the default and thus can be skipped) will visit /en and /nl variations of a website
./sitemap-generator -location=https://www.mywebsite.com -multilingual
```

These commands should generate a correctly structured `sitemap.xml` file in your current directory.
For an example of the output see the file `example-sitemap.xml` at the root of this repository.