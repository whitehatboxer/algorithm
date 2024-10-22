#include <stdio.h>
#include <stdlib.h>

struct stack
{
    int *arr;
    int len;
    int size;
};

void stack_check_size(struct stack *s, int n)
{
    if (n >= s->size)
    {
        s->arr = realloc(s->arr, 2 * s->size * sizeof(int));
        s->size = 2 * s->size;
        return;
    }
}

void stack_push(struct stack *s, int n)
{
    stack_check_size(s, s->len + 1);
    s->arr[s->len] = n;
    s->len++;
    return;
}

int stack_pop(struct stack *s)
{
    return s->arr[--s->len];
}

struct stack *init_stack(int n)
{
    struct stack *s = malloc(sizeof(struct stack));
    s->arr = malloc(sizeof(int) * n);
    s->len = 0;
    s->size = n;
    return s;
}

void get_res(int *arr, int *res, int len)
{
    struct stack *s = init_stack(len);
    int i, j;

    for (i = 0; i < len; i++)
    {
        if (s->len == 0)
        {
            stack_push(s, i);
        }
        else
        {
            j = s->len - 1;
            while (j >= 0)
            {
                if (arr[s->arr[j]] < arr[i])
                {
                    printf("i is: %d, j is : %d\n", i, s->arr[j]);
                    res[s->arr[j]] = i - s->arr[j];
                    stack_pop(s);
                    j--;
                }
                else
                {
                    break;
                }
            }
            stack_push(s, i);
        }
    }
}

int main()
{
    int arr[] = {61, 62, 63, 58, 58, 59, 64, 61};
    int arr_len = sizeof(arr) / sizeof(arr[0]);
    int *res_arr = calloc(arr_len, sizeof(arr[0]));
    get_res(arr, res_arr, arr_len);

    int i;
    printf("len is: %d\n", arr_len);
    for (i = 0; i < arr_len; i++)
    {
        printf("%d ", res_arr[i]);
    }
    printf("\n");
}
