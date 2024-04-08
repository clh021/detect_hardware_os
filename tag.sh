#!/usr/bin/env bash
# leehom Chen clh021@gmail.com

# 检查参数与tag格式
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <tag_name>"
  exit 1
fi

tag_name=$1
regex="^v[0-9]+\.[0-9]+\.[0-9]+$"

# 检查tag名称格式是否正确
if [[ ! "$tag_name" =~ $regex ]]; then
  echo "Invalid tag format. The tag should be in the form 'v0.0.1'"
  exit 1
fi

# 更新README.md文件，假设要替换第6行的内容
readme_file="README.md"
line_number=6
new_line="go get github.com/clh021/detect_hardware_os@$tag_name"

# 使用sed命令直接替换指定行内容
sed -i "${line_number}s/.*/$new_line/" "$readme_file"

# 提交更改并推送到远程仓库
git add "$readme_file"
git commit -m "Update version reference in README.md to $tag_name"
git push

# 在本地创建新的tag
git tag -a "$tag_name" -m "Creating new tag: $tag_name"

# 推送新创建的tag到远程仓库
git push origin "$tag_name"

echo "Tag '$tag_name' has been created, pushed, and updated in README.md."
