#!/bin/sh

if [ ! -d "$HOME/.oh-my-zsh" ]; then
    # 不存在oh-my-zsh
    sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
fi

ZSH_CUSTOM=$HOME/.oh-my-zsh/custom

if [ "$(grep -c "zsh-syntax-highlighting" $HOME/.zshrc)" -eq 0 ]; then
    # git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
    # git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM}/plugins/zsh-syntax-highlighting
    # sed -i "///" $HOME/.zshrc
    echo 'text here'
fi

if [ "$(grep -c "zsh-autosuggestions" $HOME/.zshrc)" -eq 0 ]; then
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM}/plugins/zsh-autosuggestions
fi
# plugins=(git zsh-syntax-highlighting zsh-autosuggestions)

# autojump
