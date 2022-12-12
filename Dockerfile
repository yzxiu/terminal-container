FROM alpine:3.13

ADD app /root
RUN chmod +x /root/app
CMD ["/root/app"]

# n build -t q946666800/newterminal_detached_mode:0.1 .
# n push q946666800/newterminal_detached_mode:0.1