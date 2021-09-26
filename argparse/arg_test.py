import argparse
parser = argparse.ArgumentParser('arg_test',
                                 description='检索指定目录下面文件并移动到新文件夹')

parser.add_argument('src', help='原始目录')
parser.add_argument('dst', help='目标目录')
# nargs='+' -t 后面至少一个值 nargs='*' -t后面至少0个参数
parser.add_argument('-t', '--type', help='文件扩展名, 如:txt md', nargs='+', required=True)

group = parser.add_mutually_exclusive_group()
group.add_argument('-s', '--silent', help='不打印移动详情', action='store_true', default=False)
group.add_argument('-v', '--verbose', help='打印移动详情', action='store_false', default=True)

parser.add_argument('-m', '--max-count', help='最大移动的数量', type=int, default=10000)

args = parser.parse_args()

print(args, '\n')

if args.verbose:
    print('正在从文件夹: %s 移动类型: %s 文件到文件夹: %s , 移动时开启打印: %s ,最大移动数量: %d' %
          (args.src, '|'.join(args.type), args.dst, args.verbose, args.max_count))
else:
    print('正在从文件夹: %s 移动类型: %s 文件到文件夹: %s , 移动时关闭打印: %s ,最大移动数量: %d' %
          (args.src,  '|'.join(args.type), args.dst, args.verbose, args.max_count))
