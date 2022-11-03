<template>
    <div class="relative z-0 w-full group">
        <label :for="name" class="form-label">{{label}}</label>
        <input
            :type="type"
            :name="name"
            :placeholder="placeholder"
            :required="required"
            :min="min"
            :max="max"
            :value="modelValue"
            @input="$emit('update:modelValue', $event.target.value)"
            class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer">
        
    </div>
</template>

<script>
export default {
    name: 'TextInput',
    props: {
        name: String,
        type: String,
        label: String,
        placeholder: String,
        required: String,
        min: String,
        max: String,
        modelValue: String,
    },
}

</script>