FROM gomicro/goose

ADD migrations/*.sql /migrations/
ADD cmd/migrations/entrypoint.sh /migrations/

#ENTRYPOINT ["/migrations/entrypoint.sh"]