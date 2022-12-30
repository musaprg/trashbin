import Head from 'next/head'
import { Stack, Heading, Container, VStack, Button, SlideFade, Box, useDisclosure } from '@chakra-ui/react'
import { LoremIpsum } from 'react-lorem-ipsum';

export default function Home() {
  const { isOpen, onToggle } = useDisclosure()
  
  return (
    <>
      <VStack>
        <Container maxW='md' bg='blue.600' color='white'>
          "md" Container
        </Container>
        <Container maxW='550px' bg='purple.600' color='white'>
          "550px" Container
        </Container>
        <Container maxW='container.sm' bg='green.400' color='#262626'>
          "container.sm" Container
        </Container>
        <Stack spacing={6}>
          <Heading as='h1' size='4xl' noOfLines={1}>
            (4xl) In love with React & Next
          </Heading>
          <Heading as='h2' size='3xl' noOfLines={1}>
            (3xl) In love with React & Next
          </Heading>
          <Heading as='h2' size='2xl'>
            (2xl) In love with React & Next
          </Heading>
          <Heading as='h2' size='xl'>
            (xl) In love with React & Next
          </Heading>
          <Heading as='h3' size='lg'>
            (lg) In love with React & Next
          </Heading>
          <Heading as='h4' size='md'>
            (md) In love with React & Next
          </Heading>
          <Heading as='h5' size='sm'>
            (sm) In love with React & Next
          </Heading>
          <Heading as='h6' size='xs'>
            (xs) In love with React & Next
          </Heading>
        </Stack>
        <Button onClick={onToggle}>Click Me</Button>
      <SlideFade in={isOpen} offsetY='20px'>
        <Box p='40px' color='white' mt='4' bg='teal.500' rounded='md' shadow='md'>
          <LoremIpsum />
        </Box>
      </SlideFade>
      </VStack>
    </>
  )
}
