<template>
  <div>
    <div v-if="this.step === 1" class="login-1">
      <img src="../../assets/login/logo.png" class="logo" />
      <div class="input_wrap">
        <!-- <span>+86</span> -->
        <input type="number" placeholder="手机号码" v-model="phone" />
      </div>
      <button class="btn" @click="getCode()">获取手机验证码</button>
      <p class="pwd" @click="$router.push('/chat')">密码登录</p>
    </div>
    <div v-if="this.step === 2" class="login-2">
      <p class="title">验证码已发送至</p>
      <p class="tel">{{phone}}</p>
      <input type="number" maxlength="4">
    </div>
  </div>
</template>
<script>
export default {
  name: "login-in",
  data() {
    return {
      phone: "222",
      step: 1
    };
  },
  methods: {
    getCode() {
      this.step = 2;
    }
  }
};
</script>
<style lang="scss">
.login-1 {
  padding-top: 0.8rem;
  text-align: center;
  font-size: 0;
  .logo {
    width: 0.89rem;
    height: 0.99rem;
  }
  .input_wrap {
    display: flex;
    margin: 0.4rem auto 0;
    padding: 0.1rem;
    width: 3.43rem;
    height: 0.4rem;
    background-color: #8e8e93;
    border-radius: 0.06rem;
    // span {
    //     font-size: 0.14rem;
    //     color: #fff;
    // }
    input {
      color: #fff;
      caret-color: #4e586e;
      &::placeholder {
        color: #4e586e;
      }
    }
  }
  .btn {
    margin-top: 0.36rem;
    width: 3.43rem;
    height: 0.4rem;
    border-radius: 0.06rem;
    background: linear-gradient(to right, #f78361, #f54b64);
    font-size: 0.15rem;
    color: #fff;
  }
  .pwd {
    margin-top: 0.3rem;
    color: #f54b64;
    font-size: 0.14rem;
    line-height: 0.2rem;
  }
}
.login-2 {
  .title {
    color: #999999;
    font-size: 0.11rem;
  }
  .tel {
    margin-top: 0.06rem;
    color: #fff;
    font-size: 0.2rem;
  }
}
</style>