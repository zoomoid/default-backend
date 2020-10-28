<template>
  <div class="container mx-auto flex flex-col justify-center h-screen" id="app">
    <img class="w-full px-4" src="./assets/undraw_Taken_re_yn20.svg" />
    <h1 class="xl:text-6xl">This route left our realms! <span class="xl:text-2xl">Make sure you requested a valid service URL.</span></h1>
    <ul>
      <a class="xl:text-2xl" href="https://alexander.bartolomey.me">https://alexander.bartolomey.me</a>
      <a class="xl:text-2xl" href="https://zoomoid.de">https://zoomoid.de</a>
    </ul>
  </div>
</template>

<script>
export default {
    name: 'App',
    components: {},
}
</script>

<style lang="sass">
#app
  -webkit-font-smoothing: antialiased
  -moz-osx-font-smoothing: grayscale
  h1
    @apply text-4xl px-4 font-semibold leading-tight
    span
      @apply block leading-normal text-xl mt-2 mb-2
  ul a
    @apply block px-4 underline
    @apply text-lg

</style>
