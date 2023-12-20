<template>
    <form @submit.prevent="submitForm">
        <input type="text" v-model="longUrl" placeholder="Enter Long URL" required />
        <input type="text" v-model="shortUrl" placeholder="Enter Short URL" />
        <button type="submit">Shorten</button>
    </form>
</template>

<script>
export default {
    data() {
        return {
            longUrl: "",
            shortUrl: ""
        };
    },
    methods: {
        submitForm() {
            // Handle form submission logic here
            // You can use an API to generate the shortened URL
            // For example:
            // axios.post("/api/shorten", { longUrl: this.longUrl, shortUrl: this.shortUrl })
            //   .then(response => {
            //     // Handle the response
            //   })
            //   .catch(error => {
            //     // Handle the error
            //   });
        }
    }
};
</script>

<style scoped>
/* Add your custom styles here */
</style>
