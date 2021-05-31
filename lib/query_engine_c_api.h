void hello(char *name);
void free_cstring(char *s);

typedef void* PrismaPtr;
typedef struct IntrospectionResult IntrospectionResult;

PrismaPtr prisma_new(char *schema);

struct IntrospectionResult {
    const char *schema;
    const char *sdl;
    const char *error;
};

IntrospectionResult* prisma_introspect (char *schema);
void free_introspection_result(IntrospectionResult* ptr);
char* prisma_get_schema (PrismaPtr ptr);
char* prisma_execute (PrismaPtr ptr,char *query);
void free_prisma (PrismaPtr ptr);
