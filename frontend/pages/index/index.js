// index.js
import api from '@/utils/api.js'
Page({
  data: {
    show: false,
    currentIndex: 0,
    duration: 0,
    playerTime: 0,
    play: false,
    volume: 5,
    songList: [],
    currentSong: null
  },
  initPlayer: function () {
    if (!this.backgroundAudioManager) {
      this.backgroundAudioManager = wx.getBackgroundAudioManager()
      this.backgroundAudioManager.volume = this.data.volume / 10
    }
    clearInterval(this.timer)
    const { currentSong } = this.data
    this.backgroundAudioManager.title = currentSong.name
    this.backgroundAudioManager.singer = currentSong.album.singer
    this.backgroundAudioManager.coverImgUrl = currentSong.album.cover
    // 注意链接不能有特殊字符，需要转换
    // this.backgroundAudioManager.src = encodeURI(this.getSrcById())
    this.backgroundAudioManager.src = encodeURI(currentSong.source)
    this.backgroundAudioManager.onCanplay(() => {
      this.backgroundAudioManager.play()
    })
    const setPlayerTime = () => {
      console.log(this.backgroundAudioManager.src);
      this.setData({
        playerTime: parseInt(this.backgroundAudioManager.currentTime),
        //duration: parseInt(this.backgroundAudioManager.duration)
      })
      return setPlayerTime
    }
    this.backgroundAudioManager.onPlay(() => {
      console.log('play');
      this.setData({
        play: true,
        duration: this.data.currentSong.duration
      })
      this.timer = setInterval(setPlayerTime(), 500)
    })
    this.backgroundAudioManager.onPause(() => {
      console.log('pause');
      clearInterval(this.timer)
      this.setData({
        play: false,
        playerTime: parseInt(this.backgroundAudioManager.currentTime)
      })
    })
    this.backgroundAudioManager.onError(e => {
      console.error('innerAudioContext.onError', e)
      wx.showToast({
        title: JSON.stringify(e),
        icon: 'error',
        duration: 2000
      })
    })
    this.backgroundAudioManager.onEnded(() => {
      clearInterval(this.timer)
      if (!this.nextPlay(null, 1)) {
        this.setData({
          play: false
        })
      }
    })
    this.backgroundAudioManager.onNext(() => {
      this.nextPlay(null, 1)
    })
    this.backgroundAudioManager.onPrev(() => {
      this.nextPlay(null, '0')
    })
    this.backgroundAudioManager.onStop(() => {
      this.setData({
        play: false,
        playerTime: 0
      })
    })
  },
  getSongList: function () {
    api.newsong().then(res => {
      const songList = res.data
      // const songList = res.result.map(item => {
      //   return {
      //     id: item.id,
      //     picUrl: item.picUrl,
      //     name: item.name,
      //     duration: parseInt(item.song.duration / 1000),
      //     artists: item.song.artists.map(a => a.name).join('/')
      //   }
      // })
      this.setData({
        songList,
        currentSong: songList[0],
        duration: songList[0].duration
      })
    })
  },
  onLoad() {
    this.getSongList()
    if (wx.createInnerAudioContext) {
      // this.initPlayer()
    } else {
      // 如果希望用户在最新版本的客户端上体验您的小程序，可以这样子提示
      wx.showModal({
        title: '提示',
        content: '当前微信版本过低，无法使用该功能，请升级到最新微信版本后重试。'
      })
    }
  },
  nextPlay: function (e, isNext) {
    const next = isNext || e.currentTarget.dataset.next
    let index = -1
    const currentIndex = this.data.currentIndex
    if (next === '0') {
      if (currentIndex === 0) return false
      index = currentIndex - 1
    } else {
      if (currentIndex === this.data.songList.length - 1) return false
      index = currentIndex + 1
    }
    this.changeSong(null, index + 1)
    return true
  },
  slider1change: function (e) {
    this.backgroundAudioManager.seek(e.detail.value)
    this.setData({
      playerTime: e.detail.value
    })
  },
  getSrcById: function () {
    const { id } = this.data.currentSong
    return 'https://music.163.com/song/media/outer/url?id=' + id + '.mp3'
  },
  handlePlay: function (e) {
    // const src = this.getSrcById(this.data.currentSong.id)
    if (this.backgroundAudioManager) {
      if (this.backgroundAudioManager.src.indexOf(this.data.currentSong.id) > 0) {
        const {
          play
        } = e.currentTarget.dataset
        if (play === '0') {
          this.backgroundAudioManager.play() // 播放
        } else {
          this.backgroundAudioManager.pause() // 暂停
        }
      } else {
        this.initPlayer()
      }
    } else {
      this.initPlayer()
    }
  },
  volumeChange: function (e) {
    const volume = e.detail.value
    this.backgroundAudioManager.volume = volume / 10
  },
  showList: function () {
    if (this.data.songList.length === 0) return
    this.setData({
      show: true
    })
  },
  changeSong: function (e, index) {
    const currentIndex = index ? index - 1 : e.currentTarget.dataset.index
    const currentSong = this.data.songList[currentIndex]
    this.setData({
      currentSong,
      currentIndex
    })
    if (this.backgroundAudioManager) {
      this.initPlayer()
    }
  }
})