import axios from 'axios';

export function getEntriesFromUser(username) {
  return new Promise(function(resolve, reject) {
    const url = `${process.env.VUE_APP_BACKEND_URL}/api/entries/${username}`;
    axios
      .get(url)
      .then(result => {
        let entries = [];
        for (const entry of result.data) {
          entries.push({
            key: entry.date,
            author: username,
            date: new Date(entry.date),
            lastModified: new Date(entry.lastModified),
            markdown: entry.markdown,
          });
        }
        // Sort newest to oldest.
        entries.sort((a, b) => b.date - a.date);
        resolve(entries);
      })
      .catch(error => {
        reject(error);
      });
  });
}