const parseGitPatch = require('parse-git-patch')


const pr = JSON.parse( fs.readFileSync('current.json', 'utf-8'))

fetch(pr.patch_url).then(res => res.text()).then(res => {
   console.log( parseGitPatch(res))
})