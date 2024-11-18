# oh-my-zsh

```bash
# 切换zsh
cat /etc/shells
echo $SHELL
chsh -s /bin/zsh
echo $SHELL

# 安装
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# 插件高亮
git clone https://gitee.com/asddfdf/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
# 插件历史提示
git clone https://gitee.com/chenweizhen/zsh-autosuggestions.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

source ~/.zshrc
# plugins=(git zsh-syntax-highlighting zsh-autosuggestions)
```
