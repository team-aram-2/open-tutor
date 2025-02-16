<script lang="ts">
  export let conversationId;
  export let messageId;
  export let originId;
  export let messageContent;
  export let messageAttachments;
  export let userId;
  export let sentOn;

  import "../../../app.css"

  const isOwnMessage = userId === originId;
</script>

<div class="message" class:own-message={isOwnMessage} class:other-message={!isOwnMessage}>
  {messageContent}
</div>

<style>
  .message {
    border-radius: 20px;
    padding: 20px 30px 20px 30px;
    margin: 0px 30px 40px 30px;
    font-size: xx-large;
    max-width: 70%;
    display:inline-block;
    word-wrap: break-word;
  }
  .own-message {
    background-color: var(--yellow-light);
    margin-left: auto;
    margin-right: 20px;
  }
  .other-message {
    background-color: var(--purple-very-light);
    margin-right: auto;
    margin-left: 20px;
  }
</style>