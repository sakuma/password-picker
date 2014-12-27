'use strict';

new Vue({
    el: '#app',
    created: function () {
      this.fetchData();
    },
    data: {
      visibilityForm: false,
      passwords: [],
      newPassword: {
        title: null,
        key: null,
        body: [
          {
            key: null,
            value: null
          },
          {
            key: null,
            value: null
          }
        ]
      }
    },
    methods: {
      fetchData: function () {
        var self = this;
        window.superagent.get('/passwords').send().end(function (res) {
          var tmp = JSON.parse(res.text);
          for (var i=0; i < tmp.length; ++i) {
            tmp[i].Body = JSON.parse(tmp[i].Body);
          }
          self.passwords = tmp;
        });
      },

      addPassword: function () {
        var self = this;
        window.superagent.post('/passwords')
          .send({
            title: self.newPassword.title,
            body: JSON.stringify(this.newPassword.body)
          })
          .end(function (res) {
            if (res.status === 200) {
              var password = JSON.parse(res.text);
              password.Body = JSON.parse(password.Body);
              console.log(password);
              self.passwords.unshift(password);
              self.visibilityForm = false;
            } else {
              console.log("failure");
              console.log(res.text);
            }
          });
        return false;
      },
      visibleAddPassword: function () {
        this.visibilityForm = true;
      }
    }
});
