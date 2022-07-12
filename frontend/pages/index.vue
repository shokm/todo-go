<template>
    <div>
        <div>{{ posts }}</div>
        <div>{{ posts.id }}</div>
        <div>{{ posts.name }}</div>
        <div>{{ posts.todo }}</div>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    async asyncData({ $axios }) {
		// 取得先のURL
		const url = "http://localhost:4000/v1/todo/2";
		// リクエスト（Get）
		const response = await $axios.$get(url);
        
		// 配列で返ってくるのでJSONにして返却
		return {
			posts: response
		};
	}
})
</script>
