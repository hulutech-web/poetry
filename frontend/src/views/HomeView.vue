<script setup lang="ts">
import {ref} from "vue";
import {GetFeather} from "../../wailsjs/go/main/App";
import {main} from "../../wailsjs/go/models";


const poetry = ref<main.Feather>({
  content: "",
  author: "",
  origin: "",
  category: "",
});

const speakTxt = () => {
  // 一边播放背景音乐，一边播放语音
  // 1、播放背景音效
  // 随机数1-6
  let randomNum = Math.floor(Math.random() * 6) + 1;
  let src = `audio${randomNum}.mp3`;
  let audio = new Audio(src);
  // 设置音效的音量
  audio.volume = 0.5;
  audio.loop = false; // 循环播放
  audio.play().catch(error => {
    console.error('播放背景音乐时出错:', error);
  });
  audio.addEventListener('ended', function () {
    console.log('背景音乐播放结束');
  })
  // 2、播放语音
  // 创建一个 SpeechSynthesisUtterance 对象
  const words = poetry.value.origin+"..."+"作者"+poetry.value.author+"..."+poetry.value.content;
  let utterance = new SpeechSynthesisUtterance(words);

  // 设置语音属性
  utterance.lang = 'zh-CN'; // 设置为中文

  // 定义一个函数来选择合适的语音
  const setVoice = () => {
    // 获取所有可用的语音
    const voices = window.speechSynthesis.getVoices();

    // 设置男声
    const maleVoice = voices.find(voice =>
        voice.lang === 'zh-CN' &&
        (voice.name.includes('Google 普通话（中国大陆）') || voice.name.includes('男'))
    );

    if (maleVoice) {
      utterance.voice = maleVoice;
    } else {
      // 如果没有找到男声，使用第一个可用的中文语音
      const anyChineseVoice = voices.find(voice => voice.lang === 'zh-CN');
      if (anyChineseVoice) {
        utterance.voice = anyChineseVoice;
        // 设置语音速度
        utterance.rate = 0.6; // 正常速度

        // 设置音调
        utterance.pitch = 0.6; // 正常音调
        console.warn('没有找到男声，使用默认的中文语音。');
      } else {
        utterance.voice = voices[0]; // 使用第一个语音
        // 设置语音速度
        utterance.rate = 0.6; // 正常速度

        // 设置音调
        utterance.pitch = 0.6; // 正常音调
        console.warn('没有找到中文语音，使用第一个可用语音。');
      }
    }

    // 使用 speechSynthesis API 播放语音
    window.speechSynthesis.speak(utterance);
  };

  // 如果语音列表已经加载，则立即设置语音，否则等待加载完成
  if (window.speechSynthesis.getVoices().length > 0) {
    setVoice();
  } else {
    window.speechSynthesis.onvoiceschanged = setVoice;
  }

  // 设置语音速度
  utterance.rate = 0.6; // 正常速度

  // 设置音调
  utterance.pitch = 0.6; // 正常音调

  // 设置语音结束后的回调函数
  utterance.onend = (event) => {
    console.log('语音播放结束');
    if(audio){
      audio.pause(); // 暂停背景音乐
    }
  };
};

const loadFeather = () => {
  let loadedPoetry = new main.Feather();
  GetFeather().then((result: main.Feather) => {
    loadedPoetry.content = result.content;
    loadedPoetry.author = result.author;
    loadedPoetry.origin = result.origin;
    loadedPoetry.category = result.category;
    poetry.value = loadedPoetry;
    speakTxt();
  })
};
const copyTxt = () => {
  // poetry.value.content;复制到剪贴板
  navigator.clipboard.writeText(poetry.value.content);
  const snackbar = document.querySelector(".my-snackbar") as any;
  if (snackbar) {
    snackbar.open = true;
  }
};
const currentThemeClass = ref("yinzhu-theme");

const changeTheme = () => {
  const themes = ["brown-theme", "enyougreen-theme", "green-theme", "purple-theme", "red-theme", "yinzhu-theme"];
  const randomIndex = Math.floor(Math.random() * themes.length);
  const randomTheme = themes[randomIndex];
  // 需要删除之前的主题样式，然后加载新的
  import(`@/theme/${currentThemeClass.value}.scss`).then(() => {
    document.documentElement.classList.remove(currentThemeClass.value);
  })
  import(`@/theme/${randomTheme}.scss`).then(() => {
    console.log(randomTheme)
    currentThemeClass.value = randomTheme;
  });
}
loadFeather();

</script>

<template>
  <div class="body">
    <div class="absolute bottom-4 right-3">
      <button @click="changeTheme" class="act-btn">
        <img src="../assets/images/color.png" class="h-4 w-4"/>
      </button>
      <button @click="loadFeather" class="act-btn">
        <img src="../assets/images/refresh.png" class="h-4 w-4"/>
      </button>
      <button @click="speakTxt" class="act-btn">
        <img src="../assets/images/voice.png" class="h-4 w-4"/>
      </button>
      <button @click="copyTxt" class="act-btn">
        <img src="../assets/images/copy.png" class="h-4 w-4"/>
      </button>
    </div>
  </div>
  <div class="home" :class="currentThemeClass">
    <div class="drag-view"></div>
    <div class="container">
      <div class="text-center text-green-200 mb-1 mt-3">
        <div style="font-family:'STFangsong'" class="text-xl font-bold">《{{ poetry.origin }}》</div>
      </div>

      <div class="content">
        <p class="txt-words text-2xl py-3">
          {{ poetry.content }}
        </p>
        <div class="txt-author text-sm">
          {{ poetry.author }}
        </div>
      </div>
      <div class="tips">
        <mdui-snackbar auto-close-delay="600" class="my-snackbar">已复制</mdui-snackbar>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
.drag-view {
  -webkit-app-region: drag;
  height: 40px; /* 你可以调整这个高度 */
  background-color: transparent; /* 保持透明或设置背景颜色 */
}

.home {
  height: 100vh;
  -webkit-user-select: none; /*webkit浏览器*/
  -moz-user-select: none; /*火狐*/
  -ms-user-select: none; /*IE10*/
  user-select: none;
  background: #212121;
}

.container {
  color: #fff;
  text-align: center;
  position: relative;
}

.act-btn {
  float: right;
  margin-right: 12px;
}

.act-btn:hover {
  transform: scale(1.2);
}
</style>
