<template>
    <div class="block w-full p-6 bg-white rounded-lg border border-gray-200 shadow-md">
        <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{{CardTitle}}</h5>
        <form @submit.prevent="handleSubmit">
            <div class="flex flex-row gap-2 mb-6">
                <TextInput name="Activity" type="text" placeholder="Whats you're going to do?" required="true"
                    v-model="formData.Activity" />
                <Button label="Add" type="submit"
                    class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-3 py-2" />
            </div>
        </form>
        <div v-for="r in result" :key="r.ID">
            <div class="flex mb-4 items-center gap-2">
                <p class="w-full text-grey-darkest line-through mr-24" v-if="r.Status">
                    {{ r.Activity }}</p>
                <p class="w-full text-grey-darkest mr-24" id="activity-status" v-else>{{ r.Activity }}</p>
                <Button label="Done"
                    class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-3 py-2"
                    v-on:click="handleDone(r.ID)" v-if="r.Status != true"/>
                <Button label="Remove"
                    class="text-white bg-red-500 hover:bg-red-600 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-3 py-2"
                    v-on:click="handelDelete(r.ID)" />
            </div>

        </div>
    </div>
</template>

<script>
    import TextInput from './forms/TextInput.vue';
    import Button from './button/Button.vue';

    export default {
        data() {
            return {
                result: {},
                formData: {
                    // id: "",
                    Activity: "",
                    Status: false,
                }
            }
        },
        props: {
            CardTitle: String,
        },
        components: {
            TextInput,
            Button,
        },
        mounted() {
            this.getData();
        },
        methods: {
            getData() {
                fetch("http://localhost:8181/todo")
                    .then(response => response.json())
                    .then(data => this.result = data)
            },

            handleSubmit() {
                fetch("http://localhost:8181/todo", {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(this.formData)
                })
                this.formData.Activity = '';
                this.getData();
            },

            handleDone(ID) {
                fetch("http://localhost:8181/todo/" + ID, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        Status: true
                    })
                });
                this.getData();

            },

            handelDelete(ID) {
                fetch("http://localhost:8181/todo/" + ID, {
                    method: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                });
                this.getData();
            }
        },
        
    }
</script>