# Evans GRPC Client

```bash

# install
curl \
  -L https://github.com/ktr0731/evans/releases/download/0.9.3/evans_linux_amd64.tar.gz \
  -o evans_linux_amd64.tar.gz
tar -xzvf evans_linux_amd64.tar.gz
rm -rf evans_linux_amd64.tar.gz
chmod +x evans && mv evans /usr/local/bin/



#################### BlogService ####################
# CreateBlog
for VALUE in {1..10}; do
  echo "{\"blog\": {\"author_id\": \"${VALUE}\", \"title\": \"Blog Title ${VALUE}\", \"content\": \"Blog Content ${VALUE}\"}}" | evans -r cli call --host localhost github.com.juliocesarscheidt.blog.BlogService.CreateBlog
done

# ReadBlog
echo '{"id": "616ae62bad7b502e79891269"}' | evans -r cli call --host localhost github.com.juliocesarscheidt.blog.BlogService.ReadBlog

# UpdateBlog
echo '{"blog": {"id": "616ae62bad7b502e79891269", "author_id": "2", "title": "Blog Title", "content": "Blog Content"}}' | evans -r cli call --host localhost github.com.juliocesarscheidt.blog.BlogService.UpdateBlog

# DeleteBlog
echo '{"id": "615e731f1061f9ca09c75f71"}' | evans -r cli call --host localhost github.com.juliocesarscheidt.blog.BlogService.DeleteBlog

# ListBlog => stream
echo '{}' | evans -r cli call --host localhost github.com.juliocesarscheidt.blog.BlogService.ListBlog



> show package

> package [NAME]

> show service

> service [SERVICE]

> show message
> desc [MESSAGE]

> call [ENDPOINT]

```
