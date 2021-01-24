<template>
  <div class="foot_wrap">
    <div class="foot_guide">
      <a
        href="javascript:;"
        class="guide_item"
        :class="{on: '/chat' === $route.path}"
        @click="goTo('/chat')"
      >
        <span class="item_icon">
        </span>
        <span>首页</span>
      </a>
      <a
        href="javascript:;"
        class="guide_item"
        :class="{on : '/contact' === $route.path}"
        @click="goTo('/contact')"
      >
        <span class="item_icon">
          <i class="iconfont icon-home"></i>
        </span>
        <span>通讯录</span>
      </a>
      <a
        href="javascript:;"
        class="guide_item"
        :class="{on : '/game' === $route.path}"
        @click="goTo('/game')"
      >
        <span class="item_icon">
          <i class="iconfont icon-home"></i>
        </span>
        <span>游戏</span>
      </a>
      <a
        href="javascript:;"
        class="guide_item"
        :class="{on : '/me' === $route.path}"
        @click="goTo('/me')"
      >
        <span class="item_icon">
          <i class="iconfont icon-home"></i>
        </span>
        <span>我</span>
      </a>
    </div>
  </div>
</template>
<script>
export default {
  name: "foot",
  mounted() {
    console.log('url', this.$route.path)
  },
  methods: {
    goTo(url) {
      this.$router.push(url);
    }
  }
};
</script>
<style lang="scss" scoped>
.foot_wrap {
  height: 0.6rem; 
}
.foot_guide {
  position: fixed;
  left: 0;
  bottom: 0;
  display: flex;
  justify-content: space-around;
  padding-top: 0.07rem;
  height: 0.6rem;
  width: 100%;
  background-color: #20242F;
  a {
    color: #4E586E;
    font-size: 0.1rem;
    text-align: center;
    .item_icon {
      display: block;
      width: 0.31rem;
      height: 0.3rem;
      background: url('../../assets/foot/message.png') no-repeat center / contain;
    }
    span {
     font-size: 0.1rem;
    }
    
    &.on {
      color: #F54B64;
      .item_icon {
        background-image: url('../../assets/foot/message_on.png');
      }
    }
  }
  a:nth-child(2) {
    .item_icon {
      display: block;
      width: 0.31rem;
      height: 0.3rem;
      background: url('../../assets/foot/contact.png') no-repeat center / contain;
    }
    &.on {
      color: #F54B64;
      .item_icon {
        background-image: url('../../assets/foot/contact_on.png');
      }
    }
  }
  a:nth-child(3) {
    .item_icon {
      display: block;
      width: 0.31rem;
      height: 0.3rem;
      background: url('../../assets/foot/game.png') no-repeat center / contain;
    }
    &.on {
      color: #F54B64;
      .item_icon {
        background-image: url('../../assets/foot/game_on.png');
      }
    }
  }
  a:nth-child(4) {
    .item_icon {
      display: block;
      width: 0.31rem;
      height: 0.3rem;
      background: url('../../assets/foot/me.png') no-repeat center / contain;
    }
    &.on {
      color: #F54B64;
      .item_icon {
        background-image: url('../../assets/foot/me_on.png');
      }
    }
  }
}
</style>