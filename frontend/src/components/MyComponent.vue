<template>
    <div>
      <input v-model="newItemName" placeholder="New item name" />
      <button @click="addItem">Add Item</button>
      <ul>
        <li v-for="item in items" :key="item.id">{{ item.name }}</li>
      </ul>
    </div>
  </template>
  
  <script>
  import { defineComponent, ref, onMounted } from 'vue';
  import { CreateItem, GetItems } from '../../wailsjs/go/backend/App';
  
  export default defineComponent({
    setup() {
      const items = ref([]);
      const newItemName = ref('');
  
      const fetchItems = async () => {
        try {
          const fetchedItems = await GetItems();
          console.log("Fetched items:", fetchedItems); // Debugging log
          items.value = fetchedItems;
        } catch (err) {
          console.error("Error fetching items:", err);
        }
      };
  
      const addItem = async () => {
        try {
          console.log("Adding item:", newItemName.value); // Debugging log
          await CreateItem(newItemName.value);
          newItemName.value = '';
          await fetchItems(); // Fetch updated items after adding new one
        } catch (err) {
          console.error("Error adding item:", err);
        }
      };
  
      onMounted(() => {
        fetchItems();
      });
  
      return { items, newItemName, addItem };
    },
  });
  </script>
  