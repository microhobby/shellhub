<template>
  <fragment>
    <v-tooltip bottom>
      <template v-slot:activator="{ on }">
        <v-icon
          v-on="on"
          @click="dialog = !dialog"
        >
          mdi-pencil
        </v-icon>
      </template>
      <span>Edit</span>
    </v-tooltip>
    <v-dialog
      v-model="dialog"
      max-width="450"
    >
      <v-card>
        <v-card-title class="headline grey lighten-2 text-center">
          Rename Device
        </v-card-title>
        <v-card-text class="caption mb-0">
          <v-text-field
            v-model="editName"
            require
            :placeholder="device.name"
            messages="Examples: (foobar, foo-bar.ba-z.qux, foo.example.com, 127.0.0.1)"
          />
        </v-card-text>
        <v-card-text
          v-if="invalid"
          class="red--text"
        >
          You entered an invalid hostname,
          it must follow the
          <a
            target="_blank"
            href="https://tools.ietf.org/html/rfc1123"
          >
            RFC1123
          </a>
          specifications
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn
            text
            @click="cancel"
          >
            Close
          </v-btn>
          <v-btn
            color="primary"
            text
            @click="check"
          >
            Rename
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </fragment>
</template>

<script>
import isValidHostname from 'is-valid-hostname';

export default {
  name: 'DeviceRename',

  props: {
    device: {
      type:Object,
      required:true
    }
  },

  data() {
    return {
      dialog: false,
      editName: '',
      invalid: false,
    };
  },
  watch:{
    dialog(out){
      !out && this.cancel();
    }
  },
  methods : {
    save() {
      this.$store.dispatch('devices/rename', {
        uid: this.device.uid,
        name: this.editName
      });
      this.dialog=false;
    },
    cancel(){
      this.dialog=false;
      this.invalid=false;
      this.editName='';
    },
    check(){
      if (isValidHostname(this.editName)){
        this.save();
        this.dialog=false;
        this.$emit('newHostname', this.editName);
        this.editName='';
      }else{
        this.invalid=true;
      }
    },
  }
};
</script>
