<template>
  <div>
    <h3>未完了</h3>
    <div v-for="n in range">
      <div v-if="posts[n - 1].status === 1" style="display: inline-flex">
        <label>
          ID: <input v-model="posts[n - 1].id" type="text" readonly />
        </label>
        <label>todo: <input v-model="posts[n - 1].todo" type="text" /></label>
        <input
          value="更新"
          type="button"
          @click="
            postTodo(posts[n - 1].id, posts[n - 1].status, posts[n - 1].todo)
          "
        />
        <input
          value="完了にする"
          type="button"
          @click="
            postTodo(posts[n - 1].id, 2, posts[n - 1].todo)
            changeStatus(posts[n - 1].id, 2)
          "
        />
      </div>
    </div>
    <h3>完了済</h3>
    <details>
      <summary>完了済タスクを見る</summary>
      <div v-for="n in range">
        <div v-if="posts[n - 1].status === 2" style="display: inline-flex">
          <label>
            ID: <input v-model="posts[n - 1].id" type="text" readonly />
          </label>
          <label>todo: <input v-model="posts[n - 1].todo" type="text" /></label>
          <input
            value="更新"
            type="button"
            @click="
              postTodo(posts[n - 1].id, posts[n - 1].status, posts[n - 1].todo)
            "
          />
          <input
            value="未完了にする"
            type="button"
            @click="
              postTodo(posts[n - 1].id, 1, posts[n - 1].todo),
                changeStatus(posts[n - 1].id, 1)
            "
          />
        </div>
      </div>
    </details>

    <h3>作成</h3>
    <div style="display: inline-flex">
      <lavel>ID: <input v-model="addIdNum" type="text" readonly /></lavel>
      <lavel>Todo: <input v-model="inputTodo" type="text" /></lavel>
      <lavel>
        Status:
        <select v-model="inputStatus">
          <option value="1">未完了</option>
          <option value="2">完了</option>
        </select>
      </lavel>
      <input
        value="送信"
        type="button"
        @click="createTodo(addIdNum, inputStatus, inputTodo)"
      />
    </div>

    <details>
      <summary>デバッグ情報</summary>
      <h3>削除（非推奨）</h3>
      <div style="display: inline-flex">
        <lavel>
          ID: {{ range }}を削除する（動作がおかしくなる場合があります）
          <input value="削除" type="button" @click="delTodo()" />
        </lavel>
      </div>
      <h3>JSON</h3>
      <div>
        {{ posts }}
      </div>
    </details>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from 'axios'

export default Vue.extend({
  async asyncData({ $axios }) {
    const resArray: any = []
    // 取得先のURL
    for (let i = 1; i < 100; i++) {
      const url = '/api/v1/todo/' + i
      // リクエスト（Get）
      const response = await $axios.$get(url)
      const resJSON = JSON.parse(JSON.stringify(response))
      if (resJSON.status === undefined) {
        break
      } else {
        resArray.push(response)
      }
    }

    // 配列で返ってくるのでJSONにして返却
    return {
      posts: JSON.parse(JSON.stringify(resArray)),
      range: resArray.length,
      addIdNum: resArray.length + 1
    }
  },
  data() {
    return {
      inputStatus: 1,
      inputTodo: ''
    }
  },
  methods: {
    createTodo(createID, createStatus, createTodo) {
      const addData = {
        id: Number(createID),
        status: Number(createStatus),
        todo: String(createTodo)
      }
      this.posts.push(addData)
      this.range++
      this.addIdNum++
      this.postTodo(createID, createStatus, createTodo)
    },
    changeStatus(id, status) {
      this.posts.find((post) => {
        if (post.id === Number(id)) {
          post.status = Number(status)
        }
      })
    },
    delTodo() {
      if (this.range !== 0) {
        this.postTodo(this.range, 0, '')
        this.range--
        this.addIdNum--
      }
    },
    postTodo(updID, updStatus, updTodo) {
      axios
        .post('/api/v1/todo/' + updID, {
          status: Number(updStatus),
          todo: String(updTodo)
        })
        .then(() => {
          // location.reload()
        })
      // .catch(() => {})
    }
  }
})
</script>
