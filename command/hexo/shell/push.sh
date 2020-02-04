#!/bin/bash
git pull

hexo clean
hexo g
hexo d

#hexo algolia

git add --all
git commit -m "auto commit"
git push origin hexo

git pull

